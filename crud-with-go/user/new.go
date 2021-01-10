package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Name string `json:"name" binding:"required"`
}

type SaveFunc func(u User) (User, error)

func SaveUser() SaveFunc {
	return func(u User) (User, error) {
		return User{}, nil
	}
}

func NewHandler(save SaveFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UserRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, req)
	}
}
