package webserver

//USAGE EXAMPLE
// func main() {
// 	r := gin.Default()
// 	db := // Your GORM database initialization
// 	userService := service.NewUserService(db)
// 	userHandler := NewUserHandler(userService)

// 	r.POST("/users", userHandler.CreateUserHandler)
// 	r.GET("/usernames", userHandler.GetAllUsernamesHandler)
// 	r.PUT("/users/:id", userHandler.UpdateUserHandler)
// 	r.POST("/users/json", userHandler.CreateUserFromJSONHandler)
// 	r.DELETE("/users/:identifier", userHandler.DeleteUserHandler)

// 	r.Run(":8080")
// }

import (
	"fmt"
	"net/http"
	"strconv"

	"fleet-monitor/backend/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// CreateUserHandler handles HTTP requests for creating a new user.
func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var request struct {
		UserName string `json:"userName"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	user, err := h.UserService.CreateUser(request.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetAllUsernamesHandler handles HTTP requests for getting all usernames.
func (h *UserHandler) GetAllUsernamesHandler(c *gin.Context) {
	usernames, err := h.UserService.GetAllUsernames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get usernames: %v", err)})
		return
	}

	c.JSON(http.StatusOK, usernames)
}

// UpdateUserHandler handles HTTP requests for updating a user.
func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	var request struct {
		UserName string `json:"userName"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	err = h.UserService.UpdateUser(uint(userID), request.UserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUserHandler handles HTTP requests for deleting a user.
func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	identifier := c.Param("identifier")

	err := h.UserService.DeleteUser(identifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
