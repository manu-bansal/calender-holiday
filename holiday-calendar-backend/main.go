package main

import (
	"fmt"
	"holiday-calendar/config"
	"holiday-calendar/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://yourfrontend.com"}, // Allow only these
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	

	routes.SetupRouter(r)

	fmt.Println("Starting server on port 8080...")
	r.Run(":8080")
}
