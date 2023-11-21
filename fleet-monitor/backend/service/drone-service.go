package service

import (
	"encoding/json"
	"fleet-monitor/backend/db"

	"gorm.io/gorm"
)

// DroneService provides methods for interacting with drones in the database.
type DroneService struct {
	db *gorm.DB
}

// NewDroneService creates a new DroneService with the given database connection.
// Example
// droneService:= service.NewDroneService(db)
func NewDroneService(db *gorm.DB) *DroneService {
	return &DroneService{db: db}
}

// CreateDrone creates a new drone with the specified details.
func (s *DroneService) CreateDrone(mavlinkID string, ownerID int) (*db.Drone, error) {
	drone := &db.Drone{
		MavlinkID: mavlinkID,
		OwnerID:   ownerID,
	}

	if err := s.db.Create(drone).Error; err != nil {
		return nil, err
	}

	return drone, nil
}

// UpdateDrone updates the drone with the given ID and sets its details.
func (s *DroneService) UpdateDrone(droneID uint, mavlinkID string, ownerID int) error {
	var drone db.Drone

	if err := s.db.First(&drone, droneID).Error; err != nil {
		return err
	}

	drone.MavlinkID = mavlinkID
	drone.OwnerID = ownerID

	if err := s.db.Save(&drone).Error; err != nil {
		return err
	}

	return nil
}

// CreateDroneFromJSON creates a new drone using JSON data.
// Example JSON request for creating a drone
// droneJSON := `{"mavlinkId": "ABC456", "ownerId": 2}`
// createDroneFromJSON(droneService, droneJSON)
func (s *DroneService) CreateDroneFromJSON(jsonStr string) (*db.Drone, error) {
	var droneData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &droneData); err != nil {
		return nil, err
	}

	mavlinkID, _ := droneData["mavlinkId"].(string)
	ownerID, _ := droneData["ownerId"].(float64)

	drone, err := s.CreateDrone(mavlinkID, int(ownerID))
	if err != nil {
		return nil, err
	}

	return drone, nil
}

// UpdateDroneRealTime updates the drone's velocity and GPS information based on JSON data.
// Input example
// velocity := Velocity{X: 2.0, Y: 1.0, Z: 0.5}
// gps := GPS{Latitude: 40.0, Longitude: -75.0}
// altitude := 100.0
func (s *DroneService) UpdateDroneRealTime(drone *db.Drone, velocity db.Velocity, gps db.GPS, altitude float64, battery int, status db.FlyingStatus) error {
	drone.Velocity = velocity
	drone.GPS = gps
	drone.Altitude = altitude
	drone.FlightStatus = status
	drone.Battery = battery

	if err := s.db.Save(drone).Error; err != nil {
		return err
	}

	return nil
}
