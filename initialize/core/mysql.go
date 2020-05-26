package core

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gomaster/common"
	"gomaster/initialize"
	"os"
)

type DBStarter struct{
	initialize.BaseStarter
}

func (db *DBStarter) Setup(ctx initialize.StarterContext) {
	conf := common.Conf()
	//数据库配置
	admin := conf.Mysql
	if db, err := gorm.Open("mysql", admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname+"?"+admin.Config); err != nil {
		common.Log().Error("MySQL启动异常", err)
		os.Exit(0)
	} else {
		common.Set_DB(db)
		common.Db().DB().SetMaxIdleConns(admin.MaxIdleConns)
		common.Db().DB().SetMaxOpenConns(admin.MaxOpenConns)
		common.Db().LogMode(admin.LogMode)
		// 全局禁用表名复数
		common.Db().SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
		//DBTables()
	}
}
