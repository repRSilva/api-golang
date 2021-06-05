package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/repRSilva/api-golang/database"
	"github.com/repRSilva/api-golang/models"
)

func ShowUser(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer",
		})

		return
	}

	db := database.GetDataBase()
	var user models.Users

	err = db.First(&user, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func ShowUsers(c *gin.Context) {

	db := database.GetDataBase()
	var users []models.Users

	err := db.Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list users: " + err.Error(),
		})

		return
	}

	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	db := database.GetDataBase()

	var user models.Users

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON",
		})

		return
	}

	err = db.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})

		return
	}

	c.JSON(200, user)
}

func EditUser(c *gin.Context) {
	db := database.GetDataBase()

	var user models.Users

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON",
		})

		return
	}

	err = db.Save(&user).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update user: " + err.Error(),
		})

		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer",
		})

		return
	}

	db := database.GetDataBase()
	err = db.Delete(models.Users{}, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete user: " + err.Error(),
		})

		return
	}

	c.Status(204)
}
