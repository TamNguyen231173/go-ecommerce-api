package middleware

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/response"
)

func AuthenMiddleware () gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrorCodeStatusUnauthorized, response.Msg[response.ErrorCodeStatusUnauthorized])

			c.Abort()
			return
		}
		c.Next()
	}
}
