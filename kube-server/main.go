package main

import (
	"kubeimooc.com/global"
	"kubeimooc.com/initiallize"
)

//项目启动入口
func main() {
	r := initiallize.Routers()
	initiallize.Viper()
	initiallize.K8S()
	panic(r.Run(global.CONF.System.Addr))
}
