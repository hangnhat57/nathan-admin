package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal-finance-app/pkg/application/user"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var input user.SignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.SignUp(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully."})
}

func (h *UserHandler) Login(c *gin.Context) {
	var input user.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) Logout(c *gin.Context) {
	// Extract user ID from context (assuming you have middleware that sets the user ID)
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	err := h.userService.Logout(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully."})
}

// Add other user-related handler methods here...
