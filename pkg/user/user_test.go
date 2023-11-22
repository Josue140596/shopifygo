package user_test

import (
	"os"
	"testing"

	"github.com/Josue140596/shopifygo/internal/db"
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

func createUserRandom(t *testing.T) {
	name := utils.RandomNames()
	email := utils.RandomEmail()
	address := utils.RandomAddress()
	password := utils.RandomPassword()
	err := userRepo.CreateUser(user.CreateUserParams{
		Username: name,
		Email:    email,
		Password: password,
		Address:  address,
	})
	require.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	createUserRandom(t)
}
