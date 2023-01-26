package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/devkishor8007/todo-api/models"
	"github.com/devkishor8007/todo-api/controllers"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// connect database
	models.Init()

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data": "hello-world"})
	})

	r.GET("/todos", controllers.AllTodos)
	r.GET("/todo/:id", controllers.FindTodo)
	r.POST("/todo", controllers.CreateTodos)
	r.PATCH("/todo/:id", controllers.UpdateTodo)
	r.DELETE("/todo/:id", controllers.DeleteTodo)

	// running port on 3007
	r.Run("0.0.0.0:3007")
}