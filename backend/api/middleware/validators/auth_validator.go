package validators

import (
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
)

func RegValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestReg *domain.Reg
		_ = c.BindJSON(&requestReg)
		if err := requestReg.Validate(); err != nil {

		}
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
