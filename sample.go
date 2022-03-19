package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//ルーターを準備
	router := gin.Default()
	//URLとハンドラを指定
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})
	//サーバーを起動
	err := router.Run(":3000")
	if err != nil {
		log.Fatal("起動に失敗しました。", err)
	}
}
