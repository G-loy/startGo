package main

import (
	"firstGin/api/user"
	"firstGin/routers"
	"fmt"
)

func main() {
	routers.Include(user.Routers)
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
