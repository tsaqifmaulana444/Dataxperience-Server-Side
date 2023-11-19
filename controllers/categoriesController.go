package controllers

import (
	"dataxperience-server-side/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexCategory(c *gin.Context) {
	var categories []models.Categories
	models.DB.Find(&categories)
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func CreateCategory(c *gin.Context) {
	var categories models.Categories

	if err := c.ShouldBindJSON(&categories); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	models.DB.Create(&categories)

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func ShowCategories(c *gin.Context) {
	var categories models.Categories
	id := c.Param("id")

	if err := models.DB.First(&categories, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message:": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func UpdateCategory(c *gin.Context) {
	var category models.Categories
	id := c.Param("id")

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	if models.DB.Model(&category).Where("id = ?", id).Updates(&category).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": "Couldn't Update Category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category Successfully Updated"})
}

func DeleteCategory(c *gin.Context) {
    id := c.Param("id")

    var category models.Categories
    if err := models.DB.First(&category, id).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Category Not Found"})
        return
    }

    if err := models.DB.Delete(&category).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed To Delete Category"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category Deleted Successfully"})
}
