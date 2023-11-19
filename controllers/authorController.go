package controllers

import (
	"dataxperience-server-side/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func CreateAuthor(c *gin.Context) {
	var author models.Authors

	if err := c.ShouldBindJSON(&author); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if isEmailTaken(author.Email) {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Email is already taken"})
		return
	}

	hashedPassword, err := hashPassword(author.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}

	author.Password = hashedPassword
	models.DB.Create(&author)

	author.Password = ""

	c.JSON(http.StatusOK, gin.H{"author": author})
}

// isEmailTaken adalah fungsi untuk memeriksa apakah email sudah digunakan

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

func isEmailTaken(email string) bool {
	var existingAuthor models.Authors
	err := models.DB.Where("email = ?", email).First(&existingAuthor).Error
	return err == nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
