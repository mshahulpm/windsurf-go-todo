package main

import (
	"log"
	"os"

	"github.com/todo-app/database"
	"github.com/todo-app/docs"
	"github.com/todo-app/handlers"
	"github.com/todo-app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Todo API
// @version         1.0
// @description     A Todo management API with user authentication and role-based access control.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Initialize database
	database.InitDB()

	// Initialize router
	r := gin.Default()

	// Swagger documentation
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize routes
	setupRoutes(r)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/api/auth/register", handlers.Register)
	r.POST("/api/auth/login", handlers.Login)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Todo routes
		protected.GET("/todos", handlers.GetTodos)
		protected.POST("/todos", handlers.CreateTodo)
		protected.GET("/todos/:id", handlers.GetTodo)
		protected.PUT("/todos/:id", handlers.UpdateTodo)
		protected.DELETE("/todos/:id", handlers.DeleteTodo)

		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.GET("/users", handlers.GetUsers)
			admin.GET("/users/:id", handlers.GetUser)
			admin.PUT("/users/:id", handlers.UpdateUser)
			admin.DELETE("/users/:id", handlers.DeleteUser)
		}
	}
}
