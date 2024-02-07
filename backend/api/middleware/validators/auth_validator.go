package validators

import "github.com/gin-gonic/gin"

func RegValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
func LoginValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
func RefreshValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
