package main

import (
	"log"
	"todo_list/internal/db"
	"todo_list/internal/handlers"
	"todo_list/internal/middleware"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
)




func main(){
	err := env.Load("../.env")
	
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	db.InitDB()

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetAllTasks)
	r.DELETE("/tasks", handlers.DeleteAllTasks)
	r.GET("/tasks/:status", handlers.GetStatusTask)
	// r.GET("/tasks/completed/", handlers.GetCompletedTasksByDate)
    r.GET("/task/:task_id", handlers.GetTaskById)
    r.PUT("/task/:task_id", handlers.UpdateTaskById)
    r.DELETE("/task/:task_id", handlers.DeleteTaskById)

	r.Run()
}	


