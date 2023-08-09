package main

import (
	"github.com/doston9471/go-crud/controllers"
	"github.com/doston9471/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main()  {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.Run()
}