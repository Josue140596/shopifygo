package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type Router struct {
	db     *UserRepository
	router *gin.Engine
}

func NewRouterUser(db *UserRepository, router *gin.Engine) *Router {
	return &Router{db, router}
}

func (r *Router) SetupRouter() {
	userGroup := r.router.Group("/users")
	{
		userGroup.POST("/create", r.createUser)
	}
}

func errorResponse(err error) gin.H {
	error := gin.H{"error": err.Error()}
	// Email error
	error = errorEmail(err)
	return error
}

func errorEmail(err error) gin.H{
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "unique_email" {
				return gin.H{"User error": "This user already exists"}
			}
		}
	}
	return gin.H{"User error": err.Error}
}