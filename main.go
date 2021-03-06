package main

import (
	"gomaster/common"
	infra "gomaster/initialize"
	"gomaster/initialize/core"
)

//启动顺序控制器
func init() {
	infra.Register(&core.PropsStarter{})
	infra.Register(&core.LogStarter{})
	infra.Register(&core.DBStarter{})
	infra.Register(&core.GinServerStarter{})
}
func main() {
	app := infra.New()
	// 程序结束前关闭数据库链接
	defer common.Db().Close()
	app.Start()
}
