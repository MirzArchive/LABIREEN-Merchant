package main

import (
	"labireen-merchant/config"
	"labireen-merchant/handlers"
	"labireen-merchant/repositories"
	"labireen-merchant/routes"
	"labireen-merchant/services"
	"labireen-merchant/utilities/mail"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(".env file loading failed")
	}

	// Initialize database connection
	db, err := config.GetDB()
	if err != nil {
		log.Fatalln("Database initialization failed")
	}

	// Auto migrate entities
	if err := config.Migrate(db); err != nil {
		log.Fatalln("Auto Migration failed")
	}

	emailService := mail.NewGmailSender(os.Getenv("EMAIL_SENDER_NAME"), os.Getenv("EMAIL_SENDER_ADDRESS"), os.Getenv("EMAIL_SENDER_PASSWORD"))
	authService := services.NewAuthService(repositories.NewMerchantRepository(db))
	MerchantService := services.NewMerchantService(repositories.NewMerchantRepository(db))

	authHandler := handlers.NewAuthHandler(authService, emailService)
	MerchantHandler := handlers.NewMerchantHandler(MerchantService)

	app := gin.Default()

	// Register auth routes
	authRoutes := routes.AuthRoutes{
		Router:      app,
		AuthHandler: authHandler,
	}
	authRoutes.Register()

	// Register Merchant routes
	MerchantRoutes := routes.MerchantRoutes{
		Router:          app,
		MerchantHandler: MerchantHandler,
	}
	MerchantRoutes.Register()

	app.Run(":" + os.Getenv("PORT"))
}
