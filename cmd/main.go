package main

import (
	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"log"
	"todo_list/internal/db"
	"todo_list/internal/handlers"
	"todo_list/internal/middleware"
)

func main() {
	err := env.Load("../env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitDB()

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetAllTasks)
	r.DELETE("/tasks", handlers.DeleteAllTasks)
	r.GET("/tasks/status/:status", handlers.GetTaskByStatus)
	r.GET("/tasks/date/:date", handlers.GetTasksByDateAndStatus)
	r.GET("/task/:task_id", handlers.GetTaskById)
	r.PUT("/task/:task_id", handlers.UpdateTaskById)
	r.DELETE("/task/:task_id", handlers.DeleteTaskById)

	r.Run(":8080")
}


