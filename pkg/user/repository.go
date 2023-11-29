package user

import (
	"errors"
	"fmt"

	"github.com/Josue140596/shopifygo/pkg/database/models"
	"github.com/Josue140596/shopifygo/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

type CreateUserParams struct {
	Username string `gorm:"size:50;not null"`
	Email    string `gorm:"size:50;not null;unique"`
	Password string `gorm:"size:50;not null"`
	Address  string `gorm:"size:100;not null"`
}

// Create a new user
func (ur *UserRepository) CreateUser(arg CreateUserParams) (models.User, error) {
	pass, _ := utils.CryptPassword(arg.Password)

	newUser := models.User{
		Username: arg.Username,
		Email:    arg.Email,
		Password: pass,
		Address:  arg.Address,
	}

	result := ur.db.Create(&newUser)
	if result.Error != nil {
		return newUser, result.Error
	}
	result.Scan(&newUser)
	return newUser, result.Error
}

// Get user by ID
func (ur *UserRepository) GetUserById(id uint) (models.User, error) {
	var user = models.User{}
	result := ur.db.First(&user, id)
	// SELECT * FROM users WHERE id = 10;
	if result.Error != nil {
		fmt.Println("failed to get user: " + result.Error.Error())
		return user, result.Error
	}
	result.Scan(&user)
	return user, result.Error
}

type UpdateUserInformationParams struct {
	UserID   uint
	Username string
	Address  string
}

// Update user
func (ur *UserRepository) UpdateUserInformation(arg UpdateUserInformationParams) (models.User, error) {
	var user models.User
	updateData := map[string]interface{}{
		"Username": arg.Username,
		"Address":  arg.Address,
	}
	// UPDATE users SET Username=?, Address=? WHERE user_id = ?;
	result := ur.db.Model(&user).Where("user_id = ?", arg.UserID).Updates(updateData)
	if result.Error != nil {
		fmt.Println("failed to update user: " + result.Error.Error())
		return user, result.Error
	}
	if result.RowsAffected == 0 {
		fmt.Println("user not found")
		return user, errors.New("user not found")
	}
	result.Scan(&user)
	return user, nil
}

// Delete user
func (ur *UserRepository) DeleteUser(id uint) error {
	var user models.User
	// DELETE FROM users WHERE user_id = ?;
	result := ur.db.Delete(&user, id)
	if result.Error != nil {
		fmt.Println("failed to delete user: " + result.Error.Error())
		return result.Error
	}
	if result.RowsAffected == 0 {
		fmt.Println("user not found")
		return errors.New("user not found")
	}

	return nil
}
