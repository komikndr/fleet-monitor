// package main

// import "fleet-monitor/backend/wails"

// func main() {
// 	wails.Run()
// }

package main

import (
	"fleet-monitor/backend/db"
	"fleet-monitor/backend/service"
	"fmt"
)

func main() {
	sql_db, err := db.OpenDB("tasks.db")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	userService := service.NewUserService(sql_db)
	droneService := service.NewDroneService(sql_db)
	taskService := service.NewTaskService(sql_db)

	userID := 1

	var user db.User
	if err := db.First(&user, userID).Error; err != nil {
		fmt.Println("Error retrieving user:", err)
		return
	}

	newUserName := "NewName"
	if err := userService.UpdateUser(user.ID, newUserName); err != nil {
		fmt.Println("Error updating user information:", err)
		return
	}

	fmt.Printf("User updated: %+v\n", user)

	droneID := 1

	var drone db.Drone
	if err := db.First(&drone, droneID).Error; err != nil {
		fmt.Println("Error retrieving drone:", err)
		return
	}

	velocity := db.Velocity{X: 2.0, Y: 1.0, Z: 0.5}
	gps := db.GPS{Latitude: 40.0, Longitude: -75.0}
	altitude := 100.0

	if err := droneService.UpdateDroneRealTime(&drone, velocity, gps, altitude); err != nil {
		fmt.Println("Error updating drone real-time information:", err)
		return
	}

	// Print the updated drone information
	fmt.Printf("Drone updated: %+v\n", drone)

	// Example usage for updating task
	taskID := 1 // Replace with the actual task ID

	// Retrieve the task from the database
	var task db.Task
	if err := db.First(&task, taskID).Error; err != nil {
		fmt.Println("Error retrieving task:", err)
		return
	}

	// Update task status
	newStatus := db.TaskStatusCompleted
	if err := taskService.UpdateTask(task.ID, newStatus); err != nil {
		fmt.Println("Error updating task status:", err)
		return
	}

	// Print the updated task information
	fmt.Printf("Task updated: %+v\n", task)
}
