package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Address  string `json:"address" binding:"required,min=8"`
}


func (server *Router) createUser(c *gin.Context) {
	 var req CreateUserRequest
	// Validations
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// Get Data
	arg := CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
	}

	_, err := server.db.CreateUser(arg)

	if err != nil {
		errorResponse(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}