package handlers

import (
	"net/http"
	"todo_list/internal/db"
	"todo_list/internal/models"
    // "time"
	"github.com/gin-gonic/gin"
)


func CreateTask(c *gin.Context){
	var newTask models.Task

    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    if err := db.DB.Create(&newTask).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, newTask)
}

func GetAllTasks(c *gin.Context){
    var TaskList []models.Task
    
    if err := db.DB.Find(&TaskList).Error; err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
        return 
    }

    c.JSON(http.StatusOK, TaskList)
}


func GetStatusTask(c *gin.Context) {
    status := c.Param("status")
    var TaskList []models.Task
    if err := db.DB.Where("status = ?", status).Find(&TaskList).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, TaskList)
}


// func GetCompletedTasksByDate(c *gin.Context) {
//     dateStr := c.Query("date")

//     date, err := time.Parse(time.RFC3339, dateStr)
//     if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
//         return
//     }

//     var tasks []models.Task
//     result := db.DB.Where("date = ? AND status = complited", date).Find(&tasks)
//     if result.Error != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//         return
//     }

//     c.JSON(http.StatusOK, tasks)
// }

func GetTaskById(c *gin.Context) {
    id := c.Param("task_id")
    var task models.Task
    if err := db.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func UpdateTaskById(c *gin.Context) {
    id := c.Param("task_id")
    var task models.Task
    if err := db.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    temp := task.ID
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.ID = temp

    if err := db.DB.Save(&task).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

func DeleteTaskById(c *gin.Context) {
    id := c.Param("task_id")
    if err := db.DB.Delete(&models.Task{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func DeleteAllTasks(c *gin.Context) {
    if err := db.DB.Where("1 = 1").Unscoped().Delete(&models.Task{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "All tasks deleted successfully"})
}