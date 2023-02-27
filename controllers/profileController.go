package controllers

import (
	"github.com/MyMap/initializers"
	"github.com/MyMap/models"
	"github.com/gin-gonic/gin"
)

// get all profiles
func ProfilesGetAll(c *gin.Context) {

	// Get all profiles from the database
	var profiles []models.Profile
	initializers.DB.Find(&profiles)

	// Return them as JSON

	c.JSON(200, profiles)

}

// create a new profile
func ProfilesCreate(c *gin.Context) {

	// Get the profile from the request
	var profile models.Profile
	c.BindJSON(&profile)

	// Save it to the database
	initializers.DB.Create(&profile)
}

// get a single profile
func ProfilesGetOne(c *gin.Context) {

	// Get the profile from the database
	var profile models.Profile

	// do a raw sql query
	initializers.DB.Raw("SELECT * FROM profiles WHERE id = ?", c.Param("id")).Scan(&profile)

	// Return it as JSON
	c.JSON(200, profile)

}

// update a profile
func ProfilesUpdate(c *gin.Context) {

	// Get the profile from the database
	var profile models.Profile
	initializers.DB.Raw("SELECT * FROM profiles WHERE id = ?", c.Param("id")).Scan(&profile)

	// Get the profile from the request
	var profileRequest models.Profile
	c.BindJSON(&profileRequest)

	// Update the profile
	profile.Username = profileRequest.Username
	profile.AvatarURL = profileRequest.AvatarURL
	profile.Bio = profileRequest.Bio
	profile.Name = profileRequest.Name
	profile.WorldMapURL = profileRequest.WorldMapURL
	profile.UpdatedAt = profileRequest.UpdatedAt

	// Save it to the database
	initializers.DB.Save(&profile)

}

// delete a profile
func ProfilesDelete(c *gin.Context) {

	// Get the profile from the database
	var profile models.Profile
	initializers.DB.Raw("SELECT * FROM profiles WHERE id = ?", c.Param("id")).Scan(&profile)

	// Delete it from the database
	initializers.DB.Delete(&profile)

}
