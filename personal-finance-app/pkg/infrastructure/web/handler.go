package web

import (
	"github.com/gin-gonic/gin"
	"personal-finance-app/pkg/application/user"
	"personal-finance-app/pkg/application/income"
	"personal-finance-app/pkg/application/expense"
)

type Handler struct {
	userHandler    *UserHandler
	incomeHandler  *IncomeHandler
	expenseHandler *ExpenseHandler
}

func NewHandler(userService user.Service, incomeService income.Service, expenseService expense.Service) *Handler {
	return &Handler{
		userHandler:    NewUserHandler(userService),
		incomeHandler:  NewIncomeHandler(incomeService),
		expenseHandler: NewExpenseHandler(expenseService),
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/signup", h.userHandler.SignUp)
			userGroup.POST("/login", h.userHandler.Login)
			userGroup.POST("/logout", h.userHandler.Logout)
			// Add other user-related routes here...
		}

		authenticated := v1.Group("/")
		authenticated.Use(h.userHandler.AuthMiddleware())
		{
			incomeGroup := authenticated.Group("/incomes")
			{
				incomeGroup.POST("/", h.incomeHandler.CreateIncome)
				incomeGroup.GET("/", h.incomeHandler.GetIncomes)
				// Add other income-related routes here...
			}

			expenseGroup := authenticated.Group("/expenses")
			{
				expenseGroup.POST("/", h.expenseHandler.CreateExpense)
				expenseGroup.GET("/", h.expenseHandler.GetExpenses)
				// Add other expense-related routes here...
			}
		}
	}
}

