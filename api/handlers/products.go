package handlers

import (
	"InjectGo-Workshop/api/service"
	"InjectGo-Workshop/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
}

type IProductController struct {
	productService service.ProductService
	userService    service.UserService
}

func NewProductController(productService service.ProductService, userService service.UserService) ProductController {
	return &IProductController{
		productService: productService,
		userService:    userService,
	}
}

func (c *IProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	//This method is called on the Gin context (ctx) and
	// is used to bind JSON data from the request body into the product variable.
	// &product is a pointer to a Go struct where you want to populate the JSON data.
	/*
		In the context of ShouldBindJSON, you pass a pointer to a struct (e.g., &product)
		because you want to populate the fields of the original struct with the data from the JSON request.
		If you were to pass just product (without the &),
		a copy of the product struct would be created within the ShouldBindJSON function,
		and any modifications made to it wouldn't affect the original product struct.
	*/
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists or not based on the provided UserID in the product
	user, err := c.userService.GetUserByID(product.UserID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	product.UserID = user.ID
	// Create the product and associate it with the user
	createdProduct, err := c.productService.CreateProduct(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdProduct)
}
