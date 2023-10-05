package routes

import (
	"InjectGo-Workshop/api/handlers"
	"InjectGo-Workshop/api/repository"
	"InjectGo-Workshop/api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func SetupUserRoutes(router *gin.RouterGroup, db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)        // Create a UserRepository instance
	userService := service.NewUserService(userRepository)     // Create a UserService instance and inject the UserRepository
	userController := handlers.NewUserController(userService) // Now you can use userService to create the UserController

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}

	// Now you can use userController to define your routes
	router.POST("/users", userController.CreateUser)
	router.GET("/users", userController.GetUsers)
}
