package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bazsup/crud-with-go/user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	initConfig()
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
		userRoutes.POST("/", user.NewHandler(user.SaveUser()))
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
