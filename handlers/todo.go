package handlers

import (
	"net/http"

	"github.com/todo-app/database"
	"github.com/todo-app/models"

	"github.com/gin-gonic/gin"
)

// @Summary     Get all todos
// @Description Get all todos for the current user (or all todos for admin)
// @Tags        todos
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Success     200 {array} models.Todo
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /todos [get]
func GetTodos(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	var todos []models.Todo
	query := database.DB

	if role != string(models.AdminRole) {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// @Summary     Create a new todo
// @Description Create a new todo for the current user
// @Tags        todos
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Param       todo body models.Todo true "Todo object"
// @Success     201 {object} models.Todo
// @Failure     400 {object} map[string]string
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /todos [post]
func CreateTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.UserID = userID.(uint)

	if err := database.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// @Summary     Get a specific todo
// @Description Get a specific todo by ID
// @Tags        todos
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Param       id path int true "Todo ID"
// @Success     200 {object} models.Todo
// @Failure     401 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Router      /todos/{id} [get]
func GetTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	id := c.Param("id")

	var todo models.Todo
	query := database.DB

	if role != string(models.AdminRole) {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// @Summary     Update a todo
// @Description Update a specific todo by ID
// @Tags        todos
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Param       id path int true "Todo ID"
// @Param       todo body models.Todo true "Todo object"
// @Success     200 {object} models.Todo
// @Failure     400 {object} map[string]string
// @Failure     401 {object} map[string]string
// @Failure     404 {object} map[string]string
// @Router      /todos/{id} [put]
func UpdateTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	id := c.Param("id")

	var todo models.Todo
	query := database.DB

	if role != string(models.AdminRole) {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// @Summary     Delete a todo
// @Description Delete a specific todo by ID
// @Tags        todos
// @Accept      json
// @Produce     json
// @Security    Bearer
// @Param       id path int true "Todo ID"
// @Success     200 {object} map[string]string
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	id := c.Param("id")

	query := database.DB

	if role != string(models.AdminRole) {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
