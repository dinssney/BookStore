package main

import (
	"BookStore/internal/db"
	"BookStore/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	db := db.Instance
	jwtKey := "super_secret_key"
	r := gin.Default()

	routes.SetupAuthRoutes(r, db, jwtKey)
	routes.SetupBooksRoutes(r, db, jwtKey)
	routes.SetupUserBooksRoutes(r, db, jwtKey)

	r.Run(":8080")
}
