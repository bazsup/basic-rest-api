package main

import (
	"database/sql"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bazsup/crud-with-go/user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	initConfig()
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	r := gin.Default()

	db, err := newDBClient(viper.GetString("db.conn.string"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", user.UserHandler(
			user.FindAllUsers(db, viper.GetString("db.table.user")),
		))
		userRoutes.POST("/", user.NewHandler(
			user.SaveUser(db, viper.GetString("db.table.user")),
			time.Now,
		))
		userRoutes.GET("/:id", user.GetOneHandler(
			user.FindById(db, viper.GetString("db.table.user")),
		))
		userRoutes.PATCH("/:id", user.UpdateHandler(
			user.UpdateUser(db, viper.GetString("db.table.user")),
			user.FindById(db, viper.GetString("db.table.user")),
			time.Now,
		))
	}

	r.Run()
}

func initConfig() {
	viper.SetDefault("db.table.user", "users")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func newDBClient(connstr string) (*sql.DB, error) {
	return sql.Open("mysql", connstr)
}
