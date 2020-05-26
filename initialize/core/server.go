package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gomaster/common"
	infra "gomaster/initialize"
)

var e *gin.Engine

type GinServerStarter struct {
	infra.BaseStarter
}

func (g *GinServerStarter) Init(ctx infra.StarterContext) {
	initGin()
}

func (g *GinServerStarter) Start(ctx infra.StarterContext) {
	address := fmt.Sprintf(":%d", common.Conf().System.Addr)
	if err := e.Run(address); err != nil {
		logrus.Fatal(err.Error())
	}
}
//阻塞
func (g *GinServerStarter) StartBlocking() bool { return true }

func initGin() {
	e = gin.Default()
	Routers(e)
}
