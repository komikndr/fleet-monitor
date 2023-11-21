package webserver

// USAGE EXAMPLE
// func main() {
// 	r := gin.Default()
// 	db := // Your GORM database initialization
// 	droneService := service.NewDroneService(db)
// 	droneHandler := NewDroneHandler(droneService)

// 	r.POST("/drones", droneHandler.CreateDroneHandler)
// 	r.GET("/drones", droneHandler.GetAllDronesHandler)
// 	r.GET("/drones/user/:userName", droneHandler.GetDronesByUserNameHandler)
// 	r.GET("/drones/taskstatus", droneHandler.GetDronesByTaskStatusHandler)
// 	r.GET("/drones/flightstatus", droneHandler.GetDronesByFlightStatusHandler)
// 	r.DELETE("/drones/:droneID", droneHandler.DeleteDroneHandler)
// 	r.POST("/drones/json", droneHandler.CreateDroneFromJSONHandler)
// 	r.PUT("/drones/:droneID/realtime", droneHandler.UpdateDroneRealTimeHandler)

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

type DroneHandler struct {
	DroneService *service.DroneService
}

func NewDroneHandler(droneService *service.DroneService) *DroneHandler {
	return &DroneHandler{DroneService: droneService}
}

// CreateDroneHandler handles HTTP requests for creating a new drone.
func (h *DroneHandler) CreateDroneHandler(c *gin.Context) {
	var request struct {
		MavlinkID string `json:"mavlinkId"`
		OwnerID   int    `json:"ownerId"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	drone, err := h.DroneService.CreateDrone(request.MavlinkID, request.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create drone: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, drone)
}

// GetAllDronesHandler handles HTTP requests for getting all drones.
func (h *DroneHandler) GetAllDronesHandler(c *gin.Context) {
	drones, err := h.DroneService.GetAllDrones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get drones: %v", err)})
		return
	}

	c.JSON(http.StatusOK, drones)
}

// GetDronesByUserNameHandler handles HTTP requests for getting drones by username.
func (h *DroneHandler) GetDronesByUserNameHandler(c *gin.Context) {
	userName := c.Param("userName")

	drones, err := h.DroneService.GetDronesByUserName(userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get drones by username: %v", err)})
		return
	}

	c.JSON(http.StatusOK, drones)
}

// GetDronesByTaskStatusHandler handles HTTP requests for getting drones by task status.
func (h *DroneHandler) GetDronesByTaskStatusHandler(c *gin.Context) {
	var request struct {
		TaskStatus string `json:"taskStatus"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	taskStatus := db.TaskStatus(request.TaskStatus)

	drones, err := h.DroneService.GetDronesByTaskStatus(taskStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get drones by task status: %v", err)})
		return
	}

	c.JSON(http.StatusOK, drones)
}

// GetDronesByFlightStatusHandler handles HTTP requests for getting drones by flight status.
func (h *DroneHandler) GetDronesByFlightStatusHandler(c *gin.Context) {
	var request struct {
		FlightStatus string `json:"flightStatus"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	flightStatus := db.FlyingStatus(request.FlightStatus)

	drones, err := h.DroneService.GetDronesByFlightStatus(flightStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get drones by flight status: %v", err)})
		return
	}

	c.JSON(http.StatusOK, drones)
}

// DeleteDroneHandler handles HTTP requests for deleting a drone by ID.
func (h *DroneHandler) DeleteDroneHandler(c *gin.Context) {
	droneIDStr := c.Param("droneID")
	droneID, err := strconv.Atoi(droneIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Drone ID"})
		return
	}

	err = h.DroneService.DeleteDroneByID(droneID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete drone: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Drone deleted successfully"})
}

// CreateDroneFromJSONHandler handles HTTP requests for creating a drone from JSON data.
// UpdateDroneRealTimeHandler handles HTTP requests for updating a drone's real-time information.
func (h *DroneHandler) UpdateDroneRealTimeHandler(c *gin.Context) {
	droneIDStr := c.Param("droneID")
	droneID, err := strconv.Atoi(droneIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Drone ID"})
		return
	}

	var request struct {
		Velocity db.Velocity     `json:"velocity"`
		GPS      db.GPS          `json:"gps"`
		Altitude float64         `json:"altitude"`
		Battery  int             `json:"battery"`
		Status   db.FlyingStatus `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	// Retrieve the drone by ID
	drone, err := h.DroneService.GetDroneByID(droneID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update drone real-time data: %v", err)})
		return
	}

	// Update the drone's real-time information
	err = h.DroneService.UpdateDroneRealTime(drone, request.Velocity, request.GPS, request.Altitude, request.Battery, request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update drone real-time data: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Drone real-time data updated successfully"})
}
