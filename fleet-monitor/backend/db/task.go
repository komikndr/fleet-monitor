package db

import "gorm.io/gorm"

type TaskStatus string

const (
	TaskStatusWaiting   TaskStatus = "waiting"
	TaskStatusOngoing   TaskStatus = "ongoing"
	TaskStatusCompleted TaskStatus = "completed"
	TaskStatusAborted   TaskStatus = "aborted"
)

// Task struct represents a task assigned to a drone.
type Task struct {
	gorm.Model
	UserID      int        `json:"userId"`
	User        User       `json:"user" gorm:"foreignKey:UserID"`
	DroneID     int        `json:"droneId"`
	Drone       Drone      `json:"drone" gorm:"foreignKey:DroneID"`
	StartLon    float64    `json:"startLon"`
	StartLat    float64    `json:"startLat"`
	EndLon      float64    `json:"endLon"`
	EndLat      float64    `json:"endLat"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}
