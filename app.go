package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AppController class
type AppController struct {
	db gorm.DB
}

// Home action
func (a *AppController) Home(c *gin.Context) {
	var allPosts []Post
	a.db.Order("ID desc").Find(&allPosts)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": allPosts,
	})
}

// Add new post action
func (a *AppController) Add(c *gin.Context) {
	var post = Post{Message: c.PostForm("Message")}
	a.db.Create(&post)
	c.Redirect(http.StatusMovedPermanently, "/")
}

// Delete a post
func (a *AppController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("ID"))
	var post = Post{ID: id}
	a.db.Delete(&post)
	c.Redirect(http.StatusMovedPermanently, "/")
}
