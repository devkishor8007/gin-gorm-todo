package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/devkishor8007/todo-api/models"
)

type CreateTodoInput struct {
	Title  string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTodoInput struct {
	Title  string `json:"title"`
	Description string `json:"description"`
}

// GET /todos
// Get all todos
func AllTodos(c *gin.Context) {
	var todos []models.Todo 
	models.DB.Find(&todos)

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

// GET /todo/:id
// Get specific todo by ID
func FindTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id= ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// POST /todo/:id
// Create a new todo
func CreateTodos(c *gin.Context) {
	// validat input 
	var input CreateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create todo
	todo := models.Todo{Title: input.Title, Description: input.Description}

	models.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// PATCH /todo/:id
// update specific todo by ID
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	var input UpdateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update todo
	updateTodoData := models.Todo{Title: input.Title, Description: input.Description}

	models.DB.Model(&todo).Updates(updateTodoData)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// DELETE /todo/:id
// delete specific todo by ID
func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
	}

	models.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"data": true})
}