package main

import (
	"fmt"
	"gin-ck/src/global"
	"gin-ck/src/route"
)

func main() {
	r := route.InitRoute()
	//r := gin.New()
	var err error
	if global.HttpVersion {
		//http服务
		if err = r.Run(global.Port); err != nil {
			fmt.Errorf("端口占用,err:%v\n", err)
		}
	} else {
		//https服务
		if err = r.RunTLS(global.Port, "https/certificate.crt", "https/private.key"); err != nil {
			fmt.Errorf("端口占用,err:%v\n", err)
		}
	}
	//if err := r.Run(global.Port); err != nil {
	//	fmt.Println("程序启动失败：", err)
	//}
	//fmt.Println("server listen 9999")
}
