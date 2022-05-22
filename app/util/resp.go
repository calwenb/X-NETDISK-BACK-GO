package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespSucce(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"code": "200", "msg": "ok", "data": data})
}
func RespFail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": "500", "msg": msg, "data": "null"})
}
func RespBadReq(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": "500", "msg": "有空参或错误", "data": "null"})
}
func RespUnAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": "401", "msg": "无权限访问", "data": "null"})
}
