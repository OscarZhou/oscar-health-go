package main

import(
	"oscar-health-go/conf"
	_ "oscar-health-go/routers" //加这行是为了激活router 里的init
)

func main(){
	app := conf.NewApp()
	app.Run()
}