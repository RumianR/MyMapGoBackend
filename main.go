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
	r.POST("/profiles", controllers.ProfilesCreate)
	r.GET("/profiles/:id", controllers.ProfilesGetOne)
	r.PUT("/profiles/:id", controllers.ProfilesUpdate)
	r.DELETE("/profiles/:id", controllers.ProfilesDelete)

	r.Run() // listen and serve on env PORT
}
