package user

import (
	"net/http"
	"time"

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

type getUserRequest struct {
	ID uint `uri:"id" binding:"required,min=1"`
}

type UserInformationRes struct {
	UserID    uint
	Username  string
	Email     string
	Address   string
	CreatedAt time.Time
}

func (server *Router) geUserById(c *gin.Context) {
	var req getUserRequest
	// Validations
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// Get Data
	user, err := server.db.GetUserById(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//Clean response
	res := UserInformationRes{
		UserID:    user.UserID,
		Username:  user.Username,
		Email:     user.Email,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, res)
}

type UpdateUserInformationReq struct {
	UserID   uint   `json:"id" binding:"required,min=1"`
	Username string `json:"username" binding:"required,min=1"`
	Address  string `json:"address" binding:"required,min=8"`
}

func (s *Router) updateUser(c *gin.Context) {
	var req UpdateUserInformationReq
	// Validations
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := UpdateUserInformationParams{
		UserID:   req.UserID,
		Username: req.Username,
		Address:  req.Address,
	}
	user, err := s.db.UpdateUserInformation(arg)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, user)
}
