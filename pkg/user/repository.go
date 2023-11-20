package user

import (
	"errors"
	"fmt"

	"github.com/Josue140596/shopifygo/pkg/database/models"
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
func (ur *UserRepository) CreateUser(arg CreateUserParams) error {
	newUser := models.User{
		Username: arg.Username,
		Email:    arg.Email,
		Password: arg.Password,
		Address:  arg.Address,
	}
	result := ur.db.Create(&newUser)
	if result.Error != nil {
		return errors.New("failed to create user: : " + result.Error.Error())
	}

	return nil
}

// Get user by ID
func (ur *UserRepository) GetUserById(id uint) (models.User, error) {
	var user = models.User{}
	result := ur.db.First(&user, id)
	// SELECT * FROM users WHERE id = 10;
	if result.Error != nil {
		fmt.Println("failed to get user: " + result.Error.Error())
	}

	return user, nil
}
