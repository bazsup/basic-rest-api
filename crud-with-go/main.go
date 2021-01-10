package main

import (
	"github.com/bazsup/crud-with-go/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users")
	userRoutes.GET("/", user.UserHandler(user.FindAllUsers()))

	r.Run()
}
