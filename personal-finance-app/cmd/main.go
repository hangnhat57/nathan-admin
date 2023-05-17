package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"personal-finance-app/pkg/application/expense"
	"personal-finance-app/pkg/application/income"
	"personal-finance-app/pkg/application/user"
	"personal-finance-app/pkg/infrastructure/persistence"
	"personal-finance-app/pkg/infrastructure/web"
)

func main() {
	// Set up Gin.
	r := gin.Default()

	// Connect to the database.
	db, err := persistence.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Initialize repositories.
	userRepo := persistence.NewUserRepository(db)
	incomeRepo := persistence.NewIncomeRepository(db)
	expenseRepo := persistence.NewExpenseRepository(db)

	// Initialize services.
	userService := user.NewUserService(userRepo)
	incomeService := income.NewIncomeService(incomeRepo)
	expenseService := expense.NewExpenseService(expenseRepo)

	// Initialize handler with all the services.
	handler := web.NewHandler(userService, incomeService, expenseService)

	// Register routes.
	r = handler.RegisterRoutes()

	// Start the server.
	if err := r.Run(); err != nil {
		log.Fatalf("server exited with error: %v", err)
	}
}
