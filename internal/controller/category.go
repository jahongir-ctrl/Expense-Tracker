package controller

import (
	"ExpenceTracker/internal/models"
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateCategoryHandler godoc
// @Summary      Create category
// @Description  Create a new category for the user
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body  object{name=string,type=string}  true  "Category info"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/categories [post]
func CreateCategoryHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	var input struct {
		Name string `json:"name" binding:"required"`
		Type string `json:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category := models.Category{
		UserID: userID,
		Name:   input.Name,
		Type:   input.Type,
	}

	if err := service.CreateCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

// GetCategoriesHandler godoc
// @Summary      Get categories
// @Description  Get all categories for the user
// @Tags         categories
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/categories [get]
func GetCategoriesHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	categories, err := service.GetCategoriesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// DeleteCategoryHandler godoc
// @Summary      Delete category
// @Description  Delete a category by ID
// @Tags         categories
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/categories/{id} [delete]
func DeleteCategoryHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	categotyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := service.DeleteCategory(categotyID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete category",
			"details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
