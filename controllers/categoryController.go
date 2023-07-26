package controllers

import (
	"errors"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/helpers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/initializers"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/models"
	"github.com/RakibSiddiquee/golang-gin-jwt-auth-crud/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"net/http"
)

// CreateCategory creates a new category
func CreateCategory(c *gin.Context) {
	// Get data from request
	var userInput struct {
		Name string `json:"name" binding:"required,min=2"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"validations": validations.FormatValidationErrors(errs),
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Name unique validation
	if !helpers.IsUniqueValue(initializers.DB, "categories", "name", userInput.Name) ||
		!helpers.IsUniqueValue(initializers.DB, "categories", "slug", slug.Make(userInput.Name)) {
		c.JSON(http.StatusConflict, gin.H{
			"validations": map[string]interface{}{
				"Name": "The name is already exist!",
			},
		})

		return
	}
	//if err := initializers.DB.Where("name = ?", userInput.Name).
	//	Or("slug = ?", slug.Make(userInput.Name)).
	//	First(&models.Category{}).Error; err == nil {
	//	c.JSON(http.StatusConflict, gin.H{
	//		"validations": map[string]interface{}{
	//			"Name": "The name is already exist!",
	//		},
	//	})
	//
	//	return
	//}

	// Create the category
	category := models.Category{
		Name: userInput.Name,
	}

	result := initializers.DB.Create(&category)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot create category",
		})

		return
	}

	// Return the category
	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

// GetCategories fetch the all categories
func GetCategories(c *gin.Context) {
	// Get the categories
	var categories []models.Category

	initializers.DB.Find(&categories)

	// Return the categories
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

// FindCategory finds the category by ID
func FindCategory(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Find the post
	var category models.Category
	result := initializers.DB.First(&category, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The record not found",
		})

		return
	}

	// Return the post
	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

// UpdateCategory updates a category
func UpdateCategory(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")

	// Get the date from request
	var userInput struct {
		Name string `json:"name" binding:"required,min=2"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"validations": validations.FormatValidationErrors(errs),
			})

			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Name unique validation
	if !helpers.IsUniqueValue(initializers.DB, "categories", "name", userInput.Name) ||
		!helpers.IsUniqueValue(initializers.DB, "categories", "slug", slug.Make(userInput.Name)) {
		c.JSON(http.StatusConflict, gin.H{
			"validations": map[string]interface{}{
				"Name": "The name is already exist!",
			},
		})

		return
	}
	//if err := initializers.DB.Where("name = ?", userInput.Name).
	//	Or("slug = ?", slug.Make(userInput.Name)).
	//	First(&models.Category{}).Error; err == nil {
	//	c.JSON(http.StatusConflict, gin.H{
	//		"validations": map[string]interface{}{
	//			"Name": "The name is already exist!",
	//		},
	//	})
	//
	//	return
	//}

	// Find the category by ID
	var category models.Category
	result := initializers.DB.First(&category, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The record not found",
		})

		return
	}

	updateCategory := models.Category{
		Name: userInput.Name,
		Slug: slug.Make(userInput.Name),
	}

	// Update the category record
	initializers.DB.Model(&category).Updates(updateCategory)

	// Return the category
	c.JSON(http.StatusOK, gin.H{
		"category": updateCategory,
	})
}

// DeleteCategory deletes a category by id
func DeleteCategory(c *gin.Context) {
	// Get the id from request
	id := c.Param("id")

	// Delete the post
	result := initializers.DB.Delete(&models.Category{}, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The category is not found",
		})

		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "The category has been deleted successfully",
	})
}

// GetTrashCategories fetch the all soft deleted categories
func GetTrashCategories(c *gin.Context) {
	// Get the categories
	var categories []models.Category

	initializers.DB.Unscoped().Find(&categories)

	// Return the categories
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func DeleteCategoryPermanent(c *gin.Context) {
	// Get the id from request
	id := c.Param("id")

	// Delete the post
	result := initializers.DB.Unscoped().Delete(&models.Category{}, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The category is not found",
		})

		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "The category has been deleted permanently",
	})
}