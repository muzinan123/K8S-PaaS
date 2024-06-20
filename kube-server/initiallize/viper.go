package initiallize

import (
	"github.com/spf13/viper" //配置文件解析库   这就是那个面试问的   yaml跟结构体做映射
	"kubeimooc.com/global"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = v.Unmarshal(&global.CONF)
	if err != nil {
		panic(err.Error())
	}
}
