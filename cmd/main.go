package main

import (
	"log"
	"os"

	"library-management-system/internal/config"
	"library-management-system/internal/handlers"
	"library-management-system/internal/middleware"
	"library-management-system/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("config.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Book{}, &models.Member{}, &models.Loan{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	r := gin.Default()
	r.Use(middleware.CORS())

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			books := protected.Group("/books")
			{
				books.GET("/", handlers.GetAllBooks)
				books.GET("/:id", handlers.GetBookByID)
				books.POST("/", handlers.CreateBook)
				books.PUT("/:id", handlers.UpdateBook)
				books.DELETE("/:id", handlers.DeleteBook)
			}

			members := protected.Group("/members")
			{
				members.GET("/", handlers.GetAllMembers)
				members.GET("/:id", handlers.GetMemberByID)
				members.POST("/", handlers.CreateMember)
				members.PUT("/:id", handlers.UpdateMember)
				members.DELETE("/:id", handlers.DeleteMember)
			}

			loans := protected.Group("/loans")
			{
				loans.GET("/", handlers.GetAllLoans)
				loans.GET("/:id", handlers.GetLoanByID)
				loans.POST("/", handlers.CreateLoan)
				loans.PUT("/:id/return", handlers.ReturnBook)
			}
		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Library Management System is running",
		})
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
