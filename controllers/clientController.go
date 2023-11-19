package controllers

import (
	"dataxperience-server-side/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateClient(c *gin.Context) {
	var client models.Clients

	if err := c.ShouldBindJSON(&client); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	models.DB.Create(&client)

	c.JSON(http.StatusOK, gin.H{"client": client})
}

func ShowClient(c *gin.Context) {
    var client models.Clients
    id := c.Param("id")

    if err := models.DB.First(&client, id).Error; err != nil {
        fmt.Println("Error fetching client:", err)
        switch err {
        case gorm.ErrRecordNotFound:
            c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message:": "Data Not Found"})
            return
        default:
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message:": err.Error()})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"client": client})
}

func UpdateClient(c *gin.Context) {
	var client models.Clients
	id := c.Param("id")

	if err := c.ShouldBindJSON(&client); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": err.Error()})
		return
	}

	if models.DB.Model(&client).Where("id = ?", id).Updates(&client).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message:": "Couldn't Update Client"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client Successfully Updated"})
}

func DeleteClient(c *gin.Context) {
    id := c.Param("id")

    var client models.Clients
    if err := models.DB.First(&client, id).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Client Not Found"})
        return
    }

    if err := models.DB.Delete(&client).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed To Delete Client"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Client Deleted Successfully"})
}
