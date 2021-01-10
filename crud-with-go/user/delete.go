package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteFunc func(int64) error

func Delete(db *sql.DB, tableName string) DeleteFunc {
	return func(id int64) error {
		sql := fmt.Sprintf(`
			DELETE FROM %s
			WHERE id = ?;
		`, tableName)

		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}

		return nil
	}
}

func DeleteHandler(delete DeleteFunc, findById FindByIdFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": `"id" must be integer`,
			})
			return
		}

		_, err = findById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": "user not found",
			})
			return
		}

		err = delete(id)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
