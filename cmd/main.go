package main

import (
	"BookStore/internal/db"
	"BookStore/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	db := db.Instance
	jwtKey := "super_secret_key"
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	routes.SetupAuthRoutes(r, db, jwtKey)
	routes.SetupBooksRoutes(r, db, jwtKey)
	routes.SetupUserBooksRoutes(r, db, jwtKey)

	r.Run(":8080")
}
