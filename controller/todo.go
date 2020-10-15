package controller

import (

	// "github.com/kzpolicy/amb-todo/service"

	"fmt"
	"strconv"
	"time"

	"local.packages/models"
	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoForm struct {
	TodoDetails models.TodoDetailSlice `json:"todos"`
	UserID      int                    `json:"user_id"`
	DateStr     string                 `json:"date"`
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

	// 日付ごとにまとめる
	td := todoDetails.
		Where(func(detail *models.TodoDetail) bool {
			// タスクが完了している過去のタスクは表示しない
			return detail.CreatedAt.Format("2006/01/02") == time.Now().Format("2006/01/02") || detail.Checked == false
		}).
		SelectTodoDetail(func(detail *models.TodoDetail) *models.TodoDetail {
			detail.DateStr = detail.CreatedAt.Format("2006/01/02")
			return detail
		}).GroupByString(func(detail *models.TodoDetail) string {
		return detail.DateStr
	})

	var todoForms []TodoForm
	for k, v := range td {
		var form TodoForm
		form.DateStr = k
		form.TodoDetails = v
		todoForms = append(todoForms, form)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todos":   todoForms,
	})
}

func CreateTodo(c *gin.Context) {

	var todoForm TodoForm
	c.BindJSON(&todoForm)

	todoService := service.NewTodoService()
	todoDetails, err := todoService.AddTodos(todoForm.TodoDetails, todoForm.UserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	}

	// 日付ごとにまとめる
	td := todoDetails.
		SelectTodoDetail(func(detail *models.TodoDetail) *models.TodoDetail {
			detail.DateStr = detail.CreatedAt.Format("2006/01/02")
			return detail
		}).GroupByString(func(detail *models.TodoDetail) string {
		return detail.DateStr
	})

	var todoForms []TodoForm
	for k, v := range td {
		var form TodoForm
		form.DateStr = k
		form.TodoDetails = v
		todoForms = append(todoForms, form)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"todos":   todoForms,
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
