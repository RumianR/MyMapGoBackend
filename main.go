package main

import (
	"github.com/MyMap/controllers"
	"github.com/MyMap/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Routes

	// Profiles
	r.GET("/profiles", controllers.ProfilesGetAll)

	r.Run() // listen and serve on env PORT
}
