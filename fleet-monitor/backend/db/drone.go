package db

import "gorm.io/gorm"

type FlyingStatus string

const (
	FlyingStatusWaiting   FlyingStatus = "damaged"
	FlyingStatusOngoing   FlyingStatus = "stable"
	FlyingStatusCompleted FlyingStatus = "offline"
	FlyingStatusAborted   FlyingStatus = "disconnected"
)

type GPS struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Velocity struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Drone struct {
	gorm.Model
	Name         string       `json:"name"`
	DroneID      int          `json:"drone_id"`
	MavlinkID    string       `json:"mavlink_id"`
	TaskID       int          `json:"task_id"`
	OwnerID      int          `json:"owner_id"`
	GPS          GPS          `json:"gps"`
	Velocity     Velocity     `json:"velocity"`
	Altitude     float64      `json:"altitude"`
	FlightStatus FlyingStatus `json:"flight_status"`
	Battery      int          `json:"battery"`
}
