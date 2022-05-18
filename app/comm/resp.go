package comm

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*func RespHSucce(data any) map[string]interface{} {
	return gin.H{"msg": "ok", "data": data}
}
func RespHFail(msg string) map[string]interface{} {
	return gin.H{"msg": msg, "data": "null"}
}*/
func RespSucce(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"msg": "ok", "data": data})
}
func RespFail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"msg": msg, "data": "null"})
}
func RespBadReq(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"msg": msg, "data": "null"})
}
