package user_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Josue140596/shopifygo/internal/db"
	"github.com/Josue140596/shopifygo/pkg/user"
)

var userRepo *user.UserRepository

func TestMain(m *testing.M) {
	db := db.NewConnection("host=localhost port=5432 user=root password=secret dbname=shopify_db sslmode=disable")
	userRepo = user.NewUserRepository(db)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestCreateUser(t *testing.T) {
	error := userRepo.CreateUser(user.CreateUserParams{
		Username: "Josue",
		Email:    "testmain2@gmail.com",
		Password: "123456",
		Address:  "Calle miguel hidalgo manzana 1 lote 1",
	})
	require.NoError(t, error)
}
