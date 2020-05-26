package core

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gomaster/common"
	_ "gomaster/docs"
	"gomaster/middleware"
	"gomaster/router"
)

//初始化总路由
func Routers(g *gin.Engine ) {
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	common.Log().Debug("use middleware logger")
	// 跨域
	g.Use(middleware.Cors())
	common.Log().Debug("use middleware cors")
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	common.Log().Debug("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := g.Group("")
	router.InitProductRouter(ApiGroup)                  // 产品路由
	common.Log().Info("router register success")
}