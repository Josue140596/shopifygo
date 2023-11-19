package main

import (
	"github.com/Josue140596/shopifygo/internal/db"
	"github.com/Josue140596/shopifygo/pkg/user"
)

func main() {
	// rootCtx := context.Background()
	dbConn := db.NewConnection("host=localhost port=5432 user=root password=secret dbname=shopify_db sslmode=disable")
	UserRepository := user.NewUserRepository(dbConn)
	UserRepository.CreateUser(user.CreateUserParams{
		Username: "Josue",
		Email:    "Josue140596@gmail.com",
		Password: "123456",
		Address:  "Street 1",
	})
}
