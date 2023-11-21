package service

import (
	"encoding/json"
	"fleet-monitor/backend/db"

	"gorm.io/gorm"
)

// TaskService provides methods for interacting with tasks in the database.
type TaskService struct {
	db *gorm.DB
}

// NewTaskService creates a new TaskService with the given database connection.
// Example
// taskService := service.NewTaskService(db)
func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{db: db}
}

// CreateTask creates a new task with the specified details and sets its status to "waiting".
func (s *TaskService) CreateTask(userID, droneID int, startLon, startLat, endLon, endLat float64, description string) (*db.Task, error) {
	task := &db.Task{
		UserID:      userID,
		DroneID:     droneID,
		StartLon:    startLon,
		StartLat:    startLat,
		EndLon:      endLon,
		EndLat:      endLat,
		Description: description,
		Status:      db.TaskStatusWaiting,
	}

	if err := s.db.Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateTask updates the task with the given ID and sets its status to the provided status.
// Example
// taskService.UpdateTask(task.ID, TaskStatusOngoing)
func (s *TaskService) UpdateTask(taskID uint, status db.TaskStatus) error {
	var task db.Task

	if err := s.db.First(&task, taskID).Error; err != nil {
		return err
	}

	task.Status = status

	if err := s.db.Save(&task).Error; err != nil {
		return err
	}

	return nil
}

// CreateTaskFromJSON creates a new task using JSON data.
// Input Example
// taskJSON := `{"userId": 1, "droneId": 1, "startLon": 10.0, "startLat": 20.0, "endLon": 15.0, "endLat": 25.0,
// "description": "Sample task"}`
// createTaskFromJSON(taskService, taskJSON)
func (s *TaskService) CreateTaskFromJSON(jsonStr string) (*db.Task, error) {
	var taskData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &taskData); err != nil {
		return nil, err
	}

	userID, _ := taskData["userId"].(float64)
	droneID, _ := taskData["droneId"].(float64)
	startLon, _ := taskData["startLon"].(float64)
	startLat, _ := taskData["startLat"].(float64)
	endLon, _ := taskData["endLon"].(float64)
	endLat, _ := taskData["endLat"].(float64)
	description, _ := taskData["description"].(string)

	task, err := s.CreateTask(int(userID), int(droneID), startLon, startLat, endLon, endLat, description)
	if err != nil {
		return nil, err
	}

	return task, nil
}
