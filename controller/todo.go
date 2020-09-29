package controller

import (

	// "github.com/kzpolicy/amb-todo/service"

	"fmt"
	"strconv"

	"local.packages/models"
	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoForm struct {
	TodoDetails []models.TodoDetail `json:"todos"`
	UserID      int                 `json:"user_id"`
}

func FindTodosByUser(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	todoService := service.NewTodoService()
	todoDetails, err := todoService.FindTodosByUser(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "ok",
		"todoDetails": todoDetails,
	})
}

func CreateTodo(c *gin.Context) {

	var todoForm TodoForm
	c.BindJSON(&todoForm)

	todoService := service.NewTodoService()
	todos, err := todoService.AddTodos(todoForm.TodoDetails, todoForm.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todos":   todos,
	})
}

func EditTodo(c *gin.Context) {

	var todos models.TodoDetailSlice
	c.BindJSON(&todos)

	todoService := service.NewTodoService()
	err := todoService.UpdateTodos(todos)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todos":   todos,
	})
}
