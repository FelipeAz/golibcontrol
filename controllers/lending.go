package controllers

import (
	"net/http"
	"time"

	"github.com/FelipeAz/golibcontrol/database/models"
	"github.com/gin-gonic/gin"
)

// GetLendings return all lendings.
func GetLendings(c *gin.Context) {
	var lendings []models.Lending

	models.DB.Find(&lendings)

	c.JSON(http.StatusOK, gin.H{"data": lendings})
}

// GetLending return one lending.
func GetLending(c *gin.Context) {
	var lending models.Lending

	if err := models.DB.Where("id = ?", c.Param("id")).First(&lending).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lending})
}

// CreateLending persist a lending into database.
func CreateLending(c *gin.Context) {
	var input models.Lending

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lending := models.Lending{BookID: input.BookID, StudentID: input.StudentID, LendingDate: time.Now()}
	models.DB.Create(&lending)

	c.JSON(http.StatusOK, gin.H{"data": lending})
}

// UpdateLending update a specific lending.
func UpdateLending(c *gin.Context) {
	var lending models.Lending
	if err := models.DB.Where("id = ?", c.Param("id")).First(&lending).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Lending
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&lending).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": lending})
}

// DeleteLending deletes a lending.
func DeleteLending(c *gin.Context) {
	var lending models.Lending

	if err := models.DB.Where("id = ?", c.Param("id")).First(&lending).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&lending)

	c.JSON(http.StatusOK, gin.H{"data": true})
}