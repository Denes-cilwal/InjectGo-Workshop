package main

import (
	"InjectGo-Workshop/api/routes"
	"InjectGo-Workshop/infrastructure"
	"InjectGo-Workshop/lib"
	"log"
)

func main() {
	// Load environment variables
	env := lib.NewEnv()

	// Create a database instance
	db, err := infrastructure.NewDatabase(env)
	if err != nil {
		log.Fatal("Failed to create a database instance:", err)
	}

	// Create a router instance
	router := infrastructure.NewRouter()
	userRoutes := router.Group("/api")
	productRoutes := router.Group("/product")

	routes.SetupUserRoutes(userRoutes, db.DB)
	routes.SetupProductRoutes(productRoutes, db.DB)
	// Start the HTTP server
	err = router.Run(":" + env.ServerPort)
	if err != nil {
		log.Fatal("Failed to start the web server:", err)
	}
}
