package main

import (
	"github.com/doston9471/go-crud/initializers"
	"github.com/doston9471/go-crud/models"
)

func init()  {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main()  {
	initializers.DB.AutoMigrate(&models.Post{})
}