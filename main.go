package main

import (
	"github.com/gin-contrib/cors"
	// "github.com/kzpolicy/amb-todo/controller"
	"github.com/kzpolicy/amb-todo/middleware"
	"local.packages/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/volatiletech/sqlboiler/boil"
)

//go:generate sqlboiler --wipe mysql

func main() {
	r := gin.Default()

	// ミドルウェア
	r.Use(middleware.RecordUaAndTime)

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	// ルーティング
	TodoRoute := r.Group("/api/v1")
	{
		v1 := TodoRoute.Group("/todo")
		{
			v1.GET("/select", controller.FindTodosByUser)
			v1.POST("/create", controller.CreateTodo)
			v1.POST("/edit", controller.EditTodo)
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	// TODO:環境変数に設定する
	r.Run(":8082")
}
