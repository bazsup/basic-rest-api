package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateFunc func(int64, User) error

func UpdateUser(db *sql.DB, tableName string) UpdateFunc {
	return func(id int64, u User) error {
		sql := fmt.Sprintf(`
			UPDATE %s
			SET name = ?, updated_at = ?
			WHERE id = ?;
		`, tableName)

		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(u.Name, u.UpdatedAt, u.ID)
		if err != nil {
			return err
		}

		return nil
	}
}

func UpdateHandler(update UpdateFunc, findById FindByIdFunc, now func() time.Time) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": `"id" must be integer`,
			})
			return
		}

		var req UserRequest
		err = c.Bind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": `"name" is not provide`,
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

		u.Name = req.Name
		u.UpdatedAt = now().UTC()
		err = update(id, u)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response(u))
	}
}
