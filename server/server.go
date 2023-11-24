package server

import (
	"github.com/Josue140596/shopifygo/pkg/user"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *user.UserRepository
	router *gin.Engine
}

func NewServer(db *user.UserRepository) *Server {
	server := &Server{db: db}
	router := gin.Default()

	// router.POST("/createAccount", server.db.CreateUser)
	// router.GET("/account/:id", server.getAccount)
	// router.GET("/accounts", server.getAccounts)
	//Gin as a service
	server.router = router
	return server
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}
