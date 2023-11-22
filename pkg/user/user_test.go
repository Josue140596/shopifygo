package user_test

import (
	"os"
	"testing"

	"github.com/Josue140596/shopifygo/internal/db"
	"github.com/Josue140596/shopifygo/pkg/database/models"
	"github.com/Josue140596/shopifygo/pkg/user"
	"github.com/Josue140596/shopifygo/pkg/utils"
	"github.com/stretchr/testify/require"
)

var userRepo *user.UserRepository

func TestMain(m *testing.M) {
	db := db.NewConnection("host=localhost port=5432 user=root password=secret dbname=shopify_db sslmode=disable")
	userRepo = user.NewUserRepository(db)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func createUserRandom(t *testing.T) models.User {
	name := utils.RandomNames()
	email := utils.RandomEmail()
	address := utils.RandomAddress()
	password := utils.RandomPassword()
	userCreated, err := userRepo.CreateUser(user.CreateUserParams{
		Username: name,
		Email:    email,
		Password: password,
		Address:  address,
	})
	require.NoError(t, err)
	require.NotEmpty(t, userCreated)
	require.Equal(t, userCreated.Username, name)
	require.Equal(t, userCreated.Email, email)
	require.Equal(t, userCreated.Address, address)
	require.NotEqual(t, userCreated.Password, password)
	require.NotEmpty(t, userCreated.Password)
	return userCreated
}

func TestCreateUser(t *testing.T) {
	createUserRandom(t)
}

func TestGetUserById(t *testing.T) {
	userCreated := createUserRandom(t)
	userGet, err := userRepo.GetUserById(userCreated.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, userGet)
	require.Equal(t, userCreated.Username, userGet.Username)
	require.Equal(t, userCreated.Email, userGet.Email)
	require.Equal(t, userCreated.Address, userGet.Address)
	require.Equal(t, userCreated.UserID, userGet.UserID)
}

func TestUpdateUserInformation(t *testing.T) {
	userCreated := createUserRandom(t)
	newName := utils.RandomNames()
	newAddress := utils.RandomAddress()
	userUpdated, err := userRepo.UpdateUserInformation(user.UpdateUserInformationParams{
		UserID:   userCreated.UserID,
		Username: newName,
		Address:  newAddress,
	})
	require.NoError(t, err)
	require.NotEmpty(t, userUpdated)
	require.Equal(t, userUpdated.Username, newName)
	require.Equal(t, userUpdated.Address, newAddress)
	require.Equal(t, userUpdated.UserID, userCreated.UserID)
	require.NotEqual(t, userUpdated.Username, userCreated.Username)
	require.NotEqual(t, userUpdated.Address, userCreated.Address)
}

func TestDeleteUser(t *testing.T) {
	userCreated := createUserRandom(t)
	err := userRepo.DeleteUser(userCreated.UserID)
	require.NoError(t, err)
	_, err = userRepo.GetUserById(userCreated.UserID)
	require.Error(t, err)
}
