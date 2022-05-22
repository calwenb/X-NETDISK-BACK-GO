package controller

import (
	"github.com/gin-gonic/gin"
	"netdisk-back/app/middle"
	"netdisk-back/app/pojo"
	"netdisk-back/app/service"
	"netdisk-back/app/util"
)

func UserRouter(e *gin.Engine) {

	userG := e.Group("/user", middle.AuthToken())
	{
		userG.GET("/get_all", getUsers)
		userG.GET("/get_all_msg", GetUsersStore)
		userG.GET("/get/:id", GetUserById)
		userG.GET("/get_like/:key", getUserLike)

		userG.DELETE("/del/:id", DelUserById)
		userG.POST("/up/:id", upUserById)
	}

}

func getUsers(c *gin.Context) {
	users := service.GetUsers()
	util.RespSucce(c, users)
}
func GetUsersStore(c *gin.Context) {
	users := service.GetStoreUsers()
	util.RespSucce(c, users)
}
func getUserLike(c *gin.Context) {
	key := c.Param("key")
	users := service.GetUserLike(key)
	util.RespSucce(c, users)
}

func GetUserById(c *gin.Context) {
	userId := c.Param("id")
	user := service.GetUserById(userId)
	util.RespSucce(c, user)
}

func DelUserById(c *gin.Context) {
	userId := c.Param("id")
	rs := service.DelUserById(userId)
	util.RespSucce(c, rs)
}

func upUserById(c *gin.Context) {
	var user pojo.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.RespBadReq(c)
		return
	}
	rs := service.UpUserById(user)
	util.RespSucce(c, rs)
}
