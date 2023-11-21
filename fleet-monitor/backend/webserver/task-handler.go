package webserver

// EXAMPLE USAGE
// func main() {
// 	r := gin.Default()
// 	db := // Your GORM database initialization
// 	taskService := service.NewTaskService(db)
// 	taskHandler := NewTaskHandler(taskService)

// 	r.POST("/tasks", taskHandler.CreateTaskHandler)
// 	r.POST("/tasks/json", taskHandler.CreateTaskFromJSONHandler)
// 	r.PUT("/tasks/:taskID", taskHandler.UpdateTaskHandler)
// 	r.GET("/tasks", taskHandler.GetAllTasksHandler)
// 	r.GET("/tasks/status", taskHandler.GetTasksByStatusHandler)

// 	r.Run(":8080")
// }

import (
	"fmt"
	"net/http"
	"strconv"

	"fleet-monitor/backend/db"
	"fleet-monitor/backend/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	TaskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{TaskService: taskService}
}

// CreateTaskHandler handles HTTP requests for creating a new task.
func (h *TaskHandler) CreateTaskHandler(c *gin.Context) {
	var request struct {
		UserID      int     `json:"userId"`
		DroneID     int     `json:"droneId"`
		StartLon    float64 `json:"startLon"`
		StartLat    float64 `json:"startLat"`
		EndLon      float64 `json:"endLon"`
		EndLat      float64 `json:"endLat"`
		Description string  `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	task, err := h.TaskService.CreateTask(request.UserID, request.DroneID, request.StartLon, request.StartLat, request.EndLon, request.EndLat, request.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create task: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// UpdateTaskHandler handles HTTP requests for updating the status of a task.
func (h *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	taskIDStr := c.Param("taskID")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
		return
	}

	var request struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	taskStatus := db.TaskStatus(request.Status)

	err = h.TaskService.UpdateTask(uint(taskID), taskStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update task: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task status updated successfully"})
}

// GetAllTasksHandler handles HTTP requests for getting all tasks.
func (h *TaskHandler) GetAllTasksHandler(c *gin.Context) {
	tasks, err := h.TaskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get tasks: %v", err)})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTasksByStatusHandler handles HTTP requests for getting tasks by status.
func (h *TaskHandler) GetTasksByStatusHandler(c *gin.Context) {
	var request struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	taskStatus := db.TaskStatus(request.Status)

	tasks, err := h.TaskService.GetTasksByStatus(taskStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get tasks by status: %v", err)})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
