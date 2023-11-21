package service

import (
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

func (s *TaskService) GetAllTasks() ([]db.Task, error) {
	var tasks []db.Task

	if err := s.db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) GetTasksByStatus(taskStatus db.TaskStatus) ([]db.Task, error) {
	var tasks []db.Task

	// Find tasks with the specified TaskStatus
	if err := s.db.Where("status = ?", taskStatus).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}
