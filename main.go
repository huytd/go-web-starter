package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
)

const (
	DB_USER = "huy"
	DB_NAME = "mdpad"
)

func connectDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", DB_USER, DB_NAME)
	db, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	return db.DB()
}

func main() {
	router := gin.Default()

	var DB = connectDB()
	var err = DB.Ping()
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		router.Static("/css", "./public/css")
		router.LoadHTMLGlob("public/*.html")

		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Hello World!!!",
			})
		})

		router.Run(":8080")
	}
}
