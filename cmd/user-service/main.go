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
		Email:    "Josue160596@gmail.com",
		Password: "123456",
		Address:  "Street 1",
	})
	// user, _ := UserRepository.GetUserById(1)
	// userToUpdate := user.UpdateUserInformationParams{
	// 	UserID:   2,
	// 	Username: "DIOS",
	// 	Address:  "Hidalgo",
	// }
	// user, _ := UserRepository.UpdateUserInformation(userToUpdate)
	// UserRepository.DeleteUser(1)
}
