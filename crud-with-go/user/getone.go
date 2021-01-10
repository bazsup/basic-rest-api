package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FindByIdFunc func(int64) (User, error)

func FindById(db *sql.DB, tableName string) FindByIdFunc {
	return func(id int64) (User, error) {
		stmt := fmt.Sprintf(`
			SELECT id, name, created_at, updated_at FROM %s WHERE id = ?
		`, tableName)

		var u User
		err := db.QueryRow(stmt, id).Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return User{}, err
		}

		return u, nil
	}
}

func GetOneHandler(findById FindByIdFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": `"id" must be integer`,
			})
			return
		}

		u, err := findById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": "user not found",
			})
			return
		}

		c.JSON(http.StatusOK, response(u))
	}
}
