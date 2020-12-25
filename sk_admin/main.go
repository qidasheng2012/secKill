package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"SecKill/sk_admin/common"
	"SecKill/sk_admin/router"
)

const (
	// server addr
	addr = "localhost:8080"
)

func main() {
	// 连接数据库
	err := common.ConnectMySQL()
	if err != nil {
		log.Printf("MySQL connect error: %v", err)
		os.Exit(1)
	}

	// 创建表
	common.CreateTable()
	// 初始化表数据
	common.InitTable()

	// 配置服务
	server := &http.Server{
		Addr:        addr,
		Handler:     common.Handler,
		ReadTimeout: 5 * time.Second,
	}

	// 注册路由
	router.RegiterRouter(common.Handler)

	// 监听服务
	err = server.ListenAndServe()
	if err != nil {
		log.Printf("start server error: %v", err)
		os.Exit(2)
	}
	log.Println("start server success")
}
