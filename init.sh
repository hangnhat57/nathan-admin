#!/bin/bash

mkdir -p personal-finance-app/cmd
touch personal-finance-app/cmd/main.go

mkdir -p personal-finance-app/pkg/application/user
touch personal-finance-app/pkg/application/user/service.go

mkdir -p personal-finance-app/pkg/application/income
touch personal-finance-app/pkg/application/income/service.go

mkdir -p personal-finance-app/pkg/application/expense
touch personal-finance-app/pkg/application/expense/service.go

mkdir -p personal-finance-app/pkg/domain/user
touch personal-finance-app/pkg/domain/user/user.go
touch personal-finance-app/pkg/domain/user/repository.go

mkdir -p personal-finance-app/pkg/domain/income
touch personal-finance-app/pkg/domain/income/income.go
touch personal-finance-app/pkg/domain/income/repository.go

mkdir -p personal-finance-app/pkg/domain/expense
touch personal-finance-app/pkg/domain/expense/expense.go
touch personal-finance-app/pkg/domain/expense/repository.go

mkdir -p personal-finance-app/pkg/infrastructure/persistence
touch personal-finance-app/pkg/infrastructure/persistence/db.go
touch personal-finance-app/pkg/infrastructure/persistence/user_repository.go
touch personal-finance-app/pkg/infrastructure/persistence/income_repository.go
touch personal-finance-app/pkg/infrastructure/persistence/expense_repository.go

mkdir -p personal-finance-app/pkg/infrastructure/web
touch personal-finance-app/pkg/infrastructure/web/handler.go
touch personal-finance-app/pkg/infrastructure/web/middleware.go

mkdir -p personal-finance-app/pkg/infrastructure/email
touch personal-finance-app/pkg/infrastructure/email/service.go

touch personal-finance-app/go.mod
touch personal-finance-app/go.sum
