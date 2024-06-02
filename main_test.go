package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo_list/internal/db"
	"todo_list/internal/handlers"
	"todo_list/internal/models"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	db.InitDB()

	router.POST("/tasks", handlers.CreateTask)
	router.GET("/tasks", handlers.GetAllTasks)
	router.GET("/tasks/:id", handlers.GetTaskById)
	router.PUT("/tasks/:id", handlers.UpdateTaskById)
	router.DELETE("/tasks/:id", handlers.DeleteTaskById)
	router.GET("/tasks/status/:status", handlers.GetTaskByStatus)
	router.GET("/tasks/date/:date", handlers.GetTasksByDateAndStatus)
	router.DELETE("/tasks", handlers.DeleteAllTasks)

	return router
}

func TestCreateTask(t *testing.T) {
	router := setupRouter()

	task := models.Task{
		Title:  "Test Task",
		Info:   "This is a test task",
		Date:   "2024-06-01",
		Status: "pending",
	}
	taskJSON, _ := json.Marshal(task)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var createdTask models.Task
	json.Unmarshal(w.Body.Bytes(), &createdTask)
	assert.Equal(t, task.Title, createdTask.Title)
}

func TestGetAllTasks(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTaskById(t *testing.T) {
	router := setupRouter()

	task := models.Task{
        Title:  "Test Task",
        Info:   "This is a test task",
        Date:   "2024-06-01",
        Status: "pending",
    }
    db.DB.Create(&task)

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", fmt.Sprintf("/tasks/%d", task.ID), nil)
    router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
        t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

    var returnedTask models.Task
    if err := json.Unmarshal(w.Body.Bytes(), &returnedTask); err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }
}

func TestUpdateTaskById(t *testing.T) {
	router := setupRouter()

	// First, create a task to update
	task := models.Task{
		Title:  "Test Task",
		Info:   "This is a test task",
		Date:   "2024-06-01",
		Status: "pending",
	}
	db.DB.Create(&task)

	// Update the task
	updatedTask := models.Task{
		Title:  "Updated Task",
		Info:   "This is an updated test task",
		Date:   "2024-06-02",
		Status: "completed",
	}
	taskJSON, _ := json.Marshal(updatedTask)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/"+fmt.Sprint(task.ID), bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var updated models.Task
	json.Unmarshal(w.Body.Bytes(), &updated)
	assert.Equal(t, updatedTask.Title, updated.Title)
}

func TestDeleteTaskById(t *testing.T) {
	router := setupRouter()

	// First, create a task to delete
	task := models.Task{
		Title:  "Test Task",
		Info:   "This is a test task",
		Date:   "2024-06-01",
		Status: "pending",
	}
	db.DB.Create(&task)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/"+fmt.Sprint(task.ID), nil)
	router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTaskByStatus(t *testing.T) {
	router := setupRouter()

	// First, create tasks to get by status
	db.DB.Exec("DELETE FROM tasks")

    // Создание задачи для теста
    task := models.Task{
        Title:  "Test Task 1",
        Info:   "This is a test task",
        Date:   "2024.06.02",
        Status: "pending",
    }
    db.DB.Create(&task)

    // Отправка запроса на получение задач по статусу
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/tasks/status/pending", nil)
    router.ServeHTTP(w, req)

    // Проверка результата
    if w.Code != http.StatusOK {
        t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

    var tasks []models.Task
    if err := json.Unmarshal(w.Body.Bytes(), &tasks); err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

    if len(tasks) != 1 {
        t.Fatalf("Expected 1 task, but got %d", len(tasks))
    }

    if tasks[0].Title != "Test Task 1" {_
        t.Fatalf("Expected task title 'Test Task 1', but got '%s'", tasks[0].Title)
    }
}

func TestGetTasksByDateAndStatus(t *testing.T) {
	router := setupRouter()

	db.DB.Exec("DELETE FROM tasks")

    // Создание задачи для теста
    task := models.Task{
        Title:  "Test Task 1",
        Info:   "This is a test task",
        Date:   "2024-06-01",
        Status: "completed",
    }
    db.DB.Create(&task)

    // Отправка запроса на получение задач по дате и статусу
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/tasks/date/2024-06-01?status=completed", nil)
    router.ServeHTTP(w, req)

    // Проверка результата
    if w.Code != http.StatusOK {
        t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
    }

    var tasks []models.Task
    if err := json.Unmarshal(w.Body.Bytes(), &tasks); err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

    if len(tasks) != 1 {
        t.Fatalf("Expected 1 task, but got %d", len(tasks))
    }

    if tasks[0].Title != "Test Task 1" {
        t.Fatalf("Expected task title 'Test Task 1', but got '%s'", tasks[0].Title)
    }
}

func TestDeleteAllTasks(t *testing.T) {
	router := setupRouter()

	// First, create tasks to delete
	task1 := models.Task{
		Title:  "Test Task 1",
		Info:   "This is a test task",
		Date:   "2024-06-01",
		Status: "pending",
	}
	task2 := models.Task{
		Title:  "Test Task 2",
		Info:   "This is another test task",
		Date:   "2024-06-02",
		Status: "completed",
	}
	db.DB.Create(&task1)
	db.DB.Create(&task2)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMain(m *testing.M) {

	os.Setenv("DB_HOST", "localhost")
    os.Setenv("DB_PORT", "5432")
    os.Setenv("DB_USER", "postgres")
    os.Setenv("DB_PASSWORD", "postgres")
    os.Setenv("DB_NAME", "postgres")
    os.Setenv("DB_SSLMODE", "disable")

	// Setup before tests if needed
	db.InitDB()

	// Run the tests
	exitVal := m.Run()

	// Teardown after tests if needed
	db.DB.Exec("DROP TABLE tasks")

	// Exit with the appropriate code
	os.Exit(exitVal)
}