package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//ルーターを準備
	router := gin.Default()
	//router.LoadHTMLGlob("static/*.html")
	//自動的なファイルを返すよう設定
	router.StaticFS("/static", http.Dir("static"))

	//ルートなら　/static/index.htmlにリダイレクト
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/static/index.html")
	})

	//router.LoadHTMLGlob("static/*.html")
	router.GET("/aa", func(c *gin.Context) {
		//c.HTML(200, "/static/aa.html", gin.H{})
		c.Redirect(200, "/static/aa.html")
	})

	//フォームの内容を受け取って挨拶する
	router.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		message := c.Query("message")
		c.Header("Content-Type", "text/html; charset=UTF-8")
		c.String(200, "<h1> Hell,"+name+"</h1> <br /> <p>"+message+"</p>")
	})

	//サーバーを起動
	err := router.Run(":3000")
	if err != nil {
		log.Fatal("起動に失敗しました。", err)
	}
}
