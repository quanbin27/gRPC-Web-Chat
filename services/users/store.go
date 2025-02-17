package users

import (
	"errors"
	"fmt"
	"github.com/quanbin27/gRPC-Web-Chat/services/types"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *UserStore {
	return &UserStore{db}
}
func (s *UserStore) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (s *UserStore) GetNameByID(id int32) (string, error) {
	var user types.User
	result := s.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("no user found with ID %d", id)
		}
		return "", result.Error
	}
	return user.Name, nil
}
func (s *UserStore) GetUsersByIDs(userIDs []int32) ([]types.User, error) {
	var users []types.User
	err := s.db.Where("id IN ?", userIDs).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *UserStore) GetUserByID(id int32) (*types.User, error) {
	var user types.User
	result := s.db.Unscoped().Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (s *UserStore) UpdateAvatar(userID int32, avatar string) error {
	err := s.db.Model(&types.User{}).Where("id = ?", userID).Update("avatar", avatar).Error
	return err
}

func (UserStore *UserStore) CreateUser(user *types.User) error {
	result := UserStore.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *UserStore) UpdateInfo(userID int32, updatedData map[string]interface{}) error {
	allowedFields := map[string]bool{
		"name":  true,
		"bio":   true,
		"email": true,
	}

	for key := range updatedData {
		if !allowedFields[key] {
			delete(updatedData, key) // Xóa các trường không hợp lệ
		}
	}

	// Cập nhật thông tin người dùng
	result := s.db.Model(&types.User{}).Where("id = ?", userID).Updates(updatedData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (s *UserStore) UpdatePassword(userID int32, password string) error {
	return s.db.Model(&types.User{}).Where("id = ?", userID).Update("password", password).Error
}
