package controllers

import (
	"dataxperience-server-side/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAuthor(c *gin.Context) {
	var author models.Authors

	if err := c.ShouldBindJSON(&author); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	models.DB.Create(&author)

	c.JSON(http.StatusOK, gin.H{"author": author})
}

func ShowAuthor(c *gin.Context) {
	var author models.Authors
	id := c.Param("id")

	if err := models.DB.First(&author, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message:": "Author Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"author": author})
}

func UpdateAuthor(c *gin.Context) {
	var author models.Authors
	id := c.Param("id")

	if err := c.ShouldBindJSON(&author); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	if models.DB.Model(&author).Where("id = ?", id).Updates(&author).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": "Couldn't Update Author"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category Successfully Updated"})
}

func DeleteAuthor(c *gin.Context) {
    id := c.Param("id")

    var author models.Authors
    if err := models.DB.First(&author, id).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Author Not Found"})
        return
    }

    if err := models.DB.Delete(&author).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed To Delete Author"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Author Deleted Successfully"})
}
