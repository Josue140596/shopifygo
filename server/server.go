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
	//Gin as a service
	server.router = router
	return server
}

func (server *Server) StartServer(address string) error {
	return server.router.Run(address)
}

func (server *Server) Router() *gin.Engine {
	return server.router
}
