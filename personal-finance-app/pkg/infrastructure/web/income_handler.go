package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal-finance-app/pkg/application/income"
)

type IncomeHandler struct {
	incomeService income.Service
}

func NewIncomeHandler(incomeService income.Service) *IncomeHandler {
	return &IncomeHandler{
		incomeService: incomeService,
	}
}

func (h *IncomeHandler) CreateIncome(c *gin.Context) {
	var input income.CreateIncomeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.incomeService.CreateIncome(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Income created successfully."})
}

func (h *IncomeHandler) GetIncomes(c *gin.Context) {
	// Extract user ID from context (assuming you have middleware that sets the user ID)
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	incomes, err := h.incomeService.GetIncomes(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, incomes)
}

// Add other income-related handler methods here...
