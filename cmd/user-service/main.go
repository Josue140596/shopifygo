package main

import (
	"github.com/Josue140596/shopifygo/internal/db"
	"github.com/Josue140596/shopifygo/pkg/user"
	"github.com/Josue140596/shopifygo/pkg/utils"
	"github.com/Josue140596/shopifygo/server"
)

func main() {
	// rootCtx := context.Background()
	// Get env variables
	config, _ := utils.LoadConfig("configs")
	// Generate DSN
	dns := utils.GenerateDSN(&config)
	// Create connection DB
	dbConn := db.NewConnection(dns)
	// Create repository (Queries)
	UserRepository := user.NewUserRepository(dbConn)
	// Create server
	server.NewServer(UserRepository).StartServer(config.ServerAddress)
}
