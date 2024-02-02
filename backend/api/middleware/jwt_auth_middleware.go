package middleware

import (
	"net/http"
	"strings"
	"team-work-space/domain"
	"team-work-space/internal/tokenutil"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		authSplit := strings.Split(authHeader, " ")
		if len(authSplit) == 2 {
			token := authSplit[1]
			authorized, err := tokenutil.IsAuthorized(token, secret)
			if authorized {
				claimsData, err := tokenutil.GetDataFromClaims(token, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
				}
				c.Set("userId", claimsData.UserId)
				c.Set("userName", claimsData.Username)
				c.Set("userRole", claimsData.RoleId)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized."})
		c.Abort()

	}
}
