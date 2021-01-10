package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// User user
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type findAllFunc func() ([]User, error)

func FindAllUsers() findAllFunc {
	return func() ([]User, error) {
		users := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
		}

		return users, nil
	}
}

func UserHandler(all findAllFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := all()

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"err": err.Error(),
			})
		}

		c.JSON(http.StatusOK, users)
	}
}
