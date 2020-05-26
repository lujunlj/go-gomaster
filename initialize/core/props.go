package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gomaster/common"
	"gomaster/initialize"
	"os"
)

const defaultConfigFile = "config.yaml"

type PropsStarter struct {
	initialize.BaseStarter
}

//初始化配置
func (p *PropsStarter) Init(ctx initialize.StarterContext) {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil{
		panic(fmt.Errorf("Fatal config file : %s \n",err))
		os.Exit(0)
	}
	v.WatchConfig()

	conf := common.Conf()
	v.OnConfigChange(func(in fsnotify.Event){
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&conf); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&conf); err!= nil{
		fmt.Println(err)
	}
	common.Set_VP(v)
	common.Set_Conf(conf)
}
