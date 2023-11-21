package service

import (
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

func (s *UserService) GetAllUsernames() ([]string, error) {
	var usernames []string

	// Select only the username column
	if err := s.db.Model(&db.User{}).Pluck("username", &usernames).Error; err != nil {
		return nil, err
	}

	return usernames, nil
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

// //////////////// THIS SECTION ISNT TESTED YET///////////////////////
func (s *UserService) DeleteUserByID(userID int) error {
	var user db.User

	if err := s.db.First(&user, userID).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUserByName deletes a user by username.
func (s *UserService) DeleteUserByName(username string) error {
	var user db.User

	if err := s.db.Where("user_name = ?", username).First(&user).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user by either username or user ID.
// If the provided value is numeric, it's considered as the user ID.
// If the provided value is a string, it's considered as the username.
// DeleteUser deletes a user by either username or user ID.
// If the provided value is not int, it's considered as the username.
func (s *UserService) DeleteUser(identifier interface{}) error {
	var user db.User

	// If identifier is not of type int, assume it's a username
	if _, ok := identifier.(int); !ok {
		// Delete by username
		if err := s.db.Where("user_name = ?", identifier).First(&user).Error; err != nil {
			return err
		}
	} else {
		// Delete by user ID
		if err := s.db.First(&user, identifier).Error; err != nil {
			return err
		}
	}

	// Delete the user
	if err := s.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
