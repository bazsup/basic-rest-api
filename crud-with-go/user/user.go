package user

import (
	"database/sql"
	"fmt"
	"log"
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

func FindAllUsers(db *sql.DB, tableName string) findAllFunc {
	return func() ([]User, error) {
		stmt := fmt.Sprintf(`
			SELECT id, name, created_at, updated_at FROM %s
		`, tableName)

		rows, err := db.Query(stmt)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		users := make([]User, 0)
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
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
