package main

import (
	"log"

	"github.com/a-rajgaria/HostelProject/repository"
	"github.com/a-rajgaria/HostelProject/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repository.DBConnect()

	// Creates a gin router with default middleware
	app  := gin.Default()


	routes.SetupRoutes(app)
	
	app.Run("localhost:8086")
}