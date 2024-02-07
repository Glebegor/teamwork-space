package validators

import (
	"net/http"
	"team-work-space/domain"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request *domain.Reg
		_ = c.ShouldBindBodyWith(&request, binding.JSON)
		if err := request.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
		}
		c.Next()
	}
}
func LoginValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request *domain.Login
		_ = c.ShouldBindBodyWith(&request, binding.JSON)
		if err := request.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
		}
		c.Next()
	}
}
func RefreshValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request *domain.Refresh
		_ = c.ShouldBindBodyWith(&request, binding.JSON)
		if err := request.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
		}
		c.Next()
	}
}
