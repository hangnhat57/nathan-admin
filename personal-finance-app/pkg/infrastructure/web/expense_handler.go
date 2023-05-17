package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal-finance-app/pkg/application/expense"
)

type ExpenseHandler struct {
	expenseService expense.Service
}

func NewExpenseHandler(expenseService expense.Service) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
	}
}

func (h *ExpenseHandler) CreateExpense(c *gin.Context) {
	var input expense.CreateExpenseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.expenseService.CreateExpense(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Expense created successfully."})
}

func (h *ExpenseHandler) GetExpenses(c *gin.Context) {
	// Extract user ID from context (assuming you have middleware that sets the user ID)
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	expenses, err := h.expenseService.GetExpenses(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

// Add other expense-related handler methods here...
