package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   int     `json:"user_id"`
	UserName string  `json:"username"`
	TaskID   int     `json:"task_id"`
	Drones   []Drone `json:"drones" gorm:"foreignKey:OwnerID"`
}
