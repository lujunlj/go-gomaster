package router

import (
	"github.com/gin-gonic/gin"
	v1 "gomaster/api/v1"
)

func InitProductRouter(Router *gin.RouterGroup) {
	router := Router.Group("product")
	{
		router.GET(":id", v1.Product) // 产品
	}
}
