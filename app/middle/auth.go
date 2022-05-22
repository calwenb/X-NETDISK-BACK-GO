package middle

import (
	"github.com/gin-gonic/gin"
	"netdisk-back/app/comm"
	"netdisk-back/app/service"
	"netdisk-back/app/util"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query(comm.TOKEN)
		if token == "" {
			c.Abort()
			util.RespUnAuth(c)
			return
		}

		if service.VerifyToken(token) {
			c.Next()
		} else {
			util.RespUnAuth(c)
			c.Abort()
		}

	}

}
