package routes

import (
	"InjectGo-Workshop/api/handlers"
	"InjectGo-Workshop/api/repository"
	"InjectGo-Workshop/api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func SetupProductRoutes(router *gin.RouterGroup, db *gorm.DB) {

	productRepository := repository.NewProductRepository(db)
	userRepository := repository.NewUserRepository(db)
	// service.NewProductService(productRepository, userRepository) // cannot use by changing order
	productService := service.NewProductService(userRepository, productRepository)
	userService := service.NewUserService(userRepository)
	productController := handlers.NewProductController(productService, userService)

	if err := productRepository.Migrate(); err != nil {
		log.Fatal("product migrate err", err)
	}

	router.POST("/add", productController.CreateProduct)

}
