package responses

import "github.com/gin-gonic/gin"

// SuccessJSON : json error response function
func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"msg": data})
}
