package infrastructure

import (
	"InjectGo-Workshop/api/responses"
	"github.com/gin-gonic/gin"
)

// Router -> Gin Router
// define your custom router struct
// This means that your custom Router struct contains all the fields and methods of the gin.Engine type,
// allowing you to extend and customize the behavior of Gin's router.
// his means that your Router struct inherits all the fields and methods of the gin.Engine type,
// as if they were defined directly within the Router struct.

type Router struct {
	*gin.Engine
}

/*
NewRouter() function serves as a Go idiom for creating and initializing
an instance of your custom Router type it returns an instance of Router with the gin.Engine embedded,
effectively initializing it.
While it's not a true constructor in the object-oriented sense,
it accomplishes the task of creating and initializing the custom type
*/

func NewRouter() Router {
	httpRouter := gin.Default()
	httpRouter.GET("/health-check", func(c *gin.Context) {
		responses.SuccessJSON(c, 200, "API Up and running 📺")
	})

	//  it will return a value of type Router,
	//  which includes all the fields and methods of gin.Engine
	return Router{
		httpRouter,
	}
}
