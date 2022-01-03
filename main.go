package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/sevipsagis/todo-api/auth"
	"github.com/sevipsagis/todo-api/todo"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("Please consider environment variable for local: %s\n", err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/token", auth.AccessToken(os.Getenv("SIGN")))

	protected := r.Group("", auth.Protect([]byte(os.Getenv("SIGN"))))
	handler := todo.NewTodoHandler(db)

	protected.POST("/todos", handler.NewTask)

	r.Run()
}