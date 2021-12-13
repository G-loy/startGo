package main

import (
	"firstGin/api/user"
	"firstGin/models"
	"firstGin/pkg/gredis"
	"firstGin/pkg/setting"
	"firstGin/routers"
	"fmt"
)

func init() {
	setting.Setup()

	models.Setup()
	fmt.Println("连接上mysql")
	err := gredis.Setup()
	if err != nil {
		fmt.Println("redis连接" + err.Error())
	}
	fmt.Println("连接上redis")
}

func main() {
	routers.Include(user.Routers)
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
