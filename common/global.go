package common

import (
	"gomaster/config"

	"github.com/jinzhu/gorm"
	oplogging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

var (
	vp   *viper.Viper
	conf *config.Server
	db   *gorm.DB
	log  *oplogging.Logger
)

func Set_VP(v *viper.Viper) {
	vp = v
}
func Vp() *viper.Viper {
	return vp
}
func Conf() *config.Server {
	return conf
}
func Set_Conf(server *config.Server) {
	conf = server
}

func Set_DB(d *gorm.DB) {
	db = d
}
func Db() *gorm.DB {
	return db
}
func Log() *oplogging.Logger {
	return log
}
func Set_LOG(l *oplogging.Logger) {
	log = l
}
