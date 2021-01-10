package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SaveFunc func(User) (int64, error)

func SaveUser(db *sql.DB, tableName string) SaveFunc {
	return func(u User) (int64, error) {
		sql := fmt.Sprintf(`
			INSERT into %s (name, created_at, updated_at)
			VALUES (?, ?, ?);
		`, tableName)

		stmt, err := db.Prepare(sql)
		if err != nil {
			return -1, err
		}

		rs, err := stmt.Exec(u.Name, u.CreatedAt, u.UpdatedAt)
		if err != nil {
			return -1, err
		}

		return rs.LastInsertId()
	}
}

func NewHandler(save SaveFunc, now func() time.Time) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UserRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		n := now()
		userForSave := req.User()
		userForSave.CreatedAt = n
		userForSave.UpdatedAt = n
		id, err := save(userForSave)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"err": err.Error(),
			})
			return
		}

		userForSave.ID = id
		c.JSON(http.StatusCreated, response(userForSave))
	}
}
