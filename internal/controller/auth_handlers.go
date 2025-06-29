package controller

import (
	"ExpenceTracker/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterHandler godoc
// @Summary      Register user
// @Description  Register a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  object{fullname=string,username=string,password=string}  true  "User info"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /register [post]
func RegisterHandler(c *gin.Context) {
	var input struct {
		Fullname string `json:"fullname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := service.RegisterUser(input.Fullname, input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User successfully registered"})

}

// LoginHandler godoc
// @Summary      Login user
// @Description  Login and get JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  object{username=string,password=string}  true  "User credentials"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      401  {object}  map[string]interface{}
// @Router       /login [post]
func LoginHandler(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := service.LoginUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "message": "Login successful"})
}

// ProfileHandler godoc
// @Summary      Get profile
// @Description  Get current user profile
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /api/profile [get]
func ProfileHandler(c *gin.Context) {
	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"username": username,
		"message":  "Profile retrieved successfully",
	})
}
