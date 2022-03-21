package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Log struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Body  string `json:"body"`
	CTime int64  `json:"ctime"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	router.GET("/aa", func(c *gin.Context) {
		getDB
	})
	// サーバーを起動しています
	router.Run(":3000")
}

func getDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("sqlite3", "db/sample.sqlite3")
	if err != nil {
		panic("接続に失敗したぽよ")
	}

	defer db.Close()
}
