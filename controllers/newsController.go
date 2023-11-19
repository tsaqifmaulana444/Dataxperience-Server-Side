package controllers

import (
	"dataxperience-server-side/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexNews(c *gin.Context) {
    var news []models.News
    models.DB.Preload("Category").Preload("Author").Find(&news)
    c.JSON(http.StatusOK, gin.H{"news": news})
}

func CreateNews(c *gin.Context) {
	var news models.News

	if err := c.ShouldBindJSON(&news); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	var category models.Categories
	if err := models.DB.First(&category, news.CategoryID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": "Category Does Not Exist"})
		return
	}

	models.DB.Create(&news)

	c.JSON(http.StatusOK, gin.H{"news": news})
}


func ShowNews(c *gin.Context) {
	var news models.News
	id := c.Param("id")

	if err := models.DB.Preload("Category").Preload("Author").First(&news, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message:": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"news": news})
}

func UpdateNews(c *gin.Context) {
	var news models.News
	id := c.Param("id")

	if err := c.ShouldBindJSON(&news); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	if models.DB.Model(&news).Where("id = ?", id).Updates(&news).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": "Couldn't Update News"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category Successfully Updated"})
}

func DeleteNews(c *gin.Context) {
    id := c.Param("id")

    var news models.News
    if err := models.DB.First(&news, id).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "News Not Found"})
        return
    }

    if err := models.DB.Delete(&news).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed To Delete News"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "News Deleted Successfully"})
}
