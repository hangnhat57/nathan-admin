#!/bin/bash

component=$1
component_lower=$(echo "$component" | tr '[:upper:]' '[:lower:]')
app_package="application"
domain_package="domain"
web_package="infrastructure/web"

# Create application/user directory and file
app_dir="pkg/$app_package/$component_lower"
app_file="$app_dir/service.go"
if [ ! -d "$app_dir" ]; then
  mkdir -p "$app_dir"
  echo "Created directory: $app_dir"
fi

if [ ! -f "$app_file" ]; then
  echo "package $app_package

import (
	// Add your imports here
)

type ${component}Service struct {
	// Add your fields here
}

func New${component}Service() *${component}Service {
	return &${component}Service{
		// Initialize your fields here
	}
}

// Add your service methods here
" > "$app_file"
  echo "Created file: $app_file"
else
  echo "File already exists: $app_file"
fi

# Create domain/user directory and file
domain_dir="pkg/$domain_package/$component_lower"
domain_file="$domain_dir/$component_lower.go"
if [ ! -d "$domain_dir" ]; then
  mkdir -p "$domain_dir"
  echo "Created directory: $domain_dir"
fi

if [ ! -f "$domain_file" ]; then
  echo "package $domain_package

type ${component} struct {
	// Add your fields here
}

// Add your domain-specific methods here
" > "$domain_file"
  echo "Created file: $domain_file"
else
  echo "File already exists: $domain_file"
fi

# Create infrastructure/web/user directory and file
web_dir="pkg/$web_package"
web_file="$web_dir/${component_lower}_handler.go"
if [ ! -d "$web_dir" ]; then
  mkdir -p "$web_dir"
  echo "Created directory: $web_dir"
fi

if [ ! -f "$web_file" ]; then
  echo "package $web_package

import (
	// Add your imports here
	"github.com/gin-gonic/gin"
	"net/http"
)

type ${component}Handler struct {
	${component_lower}Service *${app_package}.${component}Service
}

func New${component}Handler(${component_lower}Service *${app_package}.${component}Service) *${component}Handler {
	return &${component}Handler{
		${component_lower}Service: ${component_lower}Service,
	}
}

func (h *${component}Handler) Create${component} c *gin.Context) {
	// Implement your handler logic here
}

// Add other handler methods here
" > "$web_file"
  echo "Created file: $web_file"
else
  echo "File already exists: $web_file"
fi
