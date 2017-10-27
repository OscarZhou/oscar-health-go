package main

import(
	"oscar-go/conf"
	_ "oscar-go/routers" //加这行是为了激活router 里的init
)

func main(){
	conf.OApp.Run()
}