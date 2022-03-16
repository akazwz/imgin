package main

import (
	"log"
	"net/http"
	"os"

	"github.com/akazwz/imgin/initialize"
	"github.com/joho/godotenv"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func main() {
	/* 非生产环境加载 .env 配置 */
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	/* 初始化 gorm db */
	initialize.InitGormDB()

	/* 初始化路由 */
	routers := initialize.Routers()

	/* 端口地址 */
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	addr := "0.0.0.0:" + port

	s := &http.Server{
		Addr:    addr,
		Handler: routers,
	}

	id, err := gonanoid.New(7)
	if err != nil {
		log.Println(err)
	}
	log.Println(id)

	if err := s.ListenAndServe(); err != nil {
		log.Println("系统启动失败")
	}
}
