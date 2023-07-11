package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhdianrush/go-gin-product-api/config"
	"github.com/mhdianrush/go-gin-product-api/entities"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []entities.Product

	config.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func Find(c *gin.Context) {
	var product entities.Product

	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"products": product})
}

func Create(c *gin.Context) {
	var product entities.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(c *gin.Context) {
	var product entities.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if config.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "can't update products data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfull updated products data"})
}

func Delete(c *gin.Context) {
	var product entities.Product

	var Input struct {
		Id json.Number
		// means may json in int or string
	}

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := Input.Id.Int64()
	if config.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "can't delete product data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product data has been deleted"})
}
