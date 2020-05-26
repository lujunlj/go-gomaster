package core

import (
	"gomaster/common"
	"gomaster/model"
)

//注册数据库表专用
func DBTables(){
	db := common.Db()
	db.AutoMigrate(model.Product{})
	common.Log().Debug("register table success")
}