package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	router.GET("/aa", func(c *gin.Context) {
		c.HTML(200, "aa.html", gin.H{})
	})

	// サーバーを起動しています
	router.Run(":3000")
}
