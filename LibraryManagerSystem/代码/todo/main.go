package main

import (
	"fmt"
	"todo/db"
	"todo/handler"
)

func main() {
	// 初始化数据库
	db.Init()

	// 初始化web服务器
	err := handler.Start(fmt.Sprintf("%s:%s", "127.0.0.1", "8545"),"static")
	if err != nil {
		fmt.Printf("web服务错误:%v\n", err)
		return
	}

}
