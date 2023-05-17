# nathan-admin

# Personal Finance App

The Personal Finance App is a web application built using Go, Gin, GORM, and MySQL. It provides features for managing personal finances, including user management, income tracking, and expense tracking.


## Features

- User management:
  - Sign up and email verification
  - Login and logout
  - User roles and permissions (admin, VIP, normal user)
- Income tracking:
  - Create, update, and delete income records
  - Categorize income sources
- Expense tracking:
  - Create, update, and delete expense records
  - Categorize expense types
- Authentication and authorization:
  - Token-based authentication using JWT
  - Middleware for authentication and role-based authorization
- Database integration:
  - MySQL database for storing user, income, and expense data
  - GORM as the ORM library for database interactions

## Prerequisites

Make sure you have the following prerequisites installed on your system:

- Go (1.16 or higher)
- MySQL




├── cmd                     # Application entry point and setup
│   └── main.go             # Main file to start the application
├── pkg                     # Application packages
│   ├── application         # Application layer (business logic)
│   │   └── user            # User application package
│   │       └── service.go  # User service implementation
│   ├── domain              # Domain layer (domain models and interfaces)
│   │   └── user            # User domain package
│   │       ├── user.go     # User model and methods
│   │       └── repository.go # User repository interface
│   ├── infrastructure      # Infrastructure layer (database, email, web)
│   │   ├── email           # Email infrastructure package
│   │   ├── persistence     # Persistence infrastructure package
│   │   └── web             # Web infrastructure package
│   │       ├── handler.go  # Common handler logic
│   │       ├── middleware.go # Authentication middleware
│   │       ├── routes.go   # Route configuration
│   │       ├── server.go   # Web server setup and start/stop logic
│   │       └── handlers     # Specific handlers for user, income, expense
│   │           ├── user_handler.go
│   │           ├── income_handler.go
│   │           └── expense_handler.go
└── README.md               # Project README file
