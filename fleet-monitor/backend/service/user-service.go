package service

import (
	"encoding/json"
	"fleet-monitor/backend/db"

	"gorm.io/gorm"
)

// UserService provides methods for interacting with users in the database.
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new UserService with the given database connection.
// Example
// userService := service.NewUserService(db)
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// CreateUser creates a new user with the specified details.
func (s *UserService) CreateUser(userName string) (*db.User, error) {
	user := &db.User{
		UserName: userName,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates the user with the given ID and sets its details.
func (s *UserService) UpdateUser(userID uint, userName string) error {
	var user db.User

	if err := s.db.First(&user, userID).Error; err != nil {
		return err
	}

	user.UserName = userName

	if err := s.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

// CreateUserFromJSON creates a new user using JSON data.
// Example JSON request for creating a user
// userJSON := `{"userName": "Alice"}`
// createUserFromJSON(userService, userJSON)
func (s *UserService) CreateUserFromJSON(jsonStr string) (*db.User, error) {
	var userData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &userData); err != nil {
		return nil, err
	}

	userName, _ := userData["userName"].(string)

	user, err := s.CreateUser(userName)
	if err != nil {
		return nil, err
	}

	return user, nil
}
