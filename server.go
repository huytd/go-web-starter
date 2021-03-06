package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Server class
type Server struct {
	ssl []string
}

// Start server, all route come here
func (s *Server) Start(DBUSER string, DBPWD string, DBNAME string) {
	var connString = fmt.Sprintf("user=%s dbname=%s sslmode=disable", DBUSER, DBNAME)
	DB, _ := gorm.Open("postgres", connString)

	if ok := DB.HasTable("posts"); ok {
		fmt.Println("Checking DB.Post OK!")
	} else {
		DB.CreateTable(&Post{})
	}

	var router = gin.Default()
	var app = &AppController{db: DB}

	router.Static("/css", "./public/css")
	router.LoadHTMLGlob("public/*.html")

	// All routing here
	router.GET("/", app.Home)
	router.POST("/add", app.Add)
	router.POST("/delete", app.Delete)

	// Start HTTP server
	go func() {
		http.ListenAndServe(":8000", router)
	}()

	// Start HTTPS server as well
	if len(s.ssl) > 0 {
		http.ListenAndServeTLS(":8001", s.ssl[0], s.ssl[1], router)
	}
}
