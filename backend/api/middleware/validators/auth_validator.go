package validators

import (
	"net/http"
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestReg *domain.Reg
		_ = c.ShouldBindBodyWith(&requestReg, binding.JSON)
		if err := requestReg.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Bad info"})
			c.Abort()
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
