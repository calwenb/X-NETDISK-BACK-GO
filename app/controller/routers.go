package controller

//路由注册中心
import (
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{
	LoginRouter,
	UserRouter,
}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	app := gin.New()
	//初始化全局中间件
	//app.Use(middleware.Middleware())
	for _, opt := range options {
		opt(app)
	}
	return app
}
