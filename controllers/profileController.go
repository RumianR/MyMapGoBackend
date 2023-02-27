package controllers

import (
	"github.com/MyMap/initializers"
	"github.com/MyMap/models"
	"github.com/gin-gonic/gin"
)

func ProfilesGetAll(c *gin.Context) {

	// Get all profiles from the database
	var profiles []models.Profile
	initializers.DB.Find(&profiles)

	// Return them as JSON

	c.JSON(200, gin.H{
		"profiles": profiles,
	})

}
