package api

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (server *Server) CartUserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		cartId, err := strconv.Atoi(c.Param("cartId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse("Bad Request", errors.New("invalid params id")))
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)

		userID := uint(userData["id"].(float64))

		cart, err := server.store.GetOneCartById(context.Background(), int64(cartId))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse("Data Not Found", errors.New("cart data doens't exist")))
			return
		}

		if uint(cart.UserID) != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse("Internal Server Error", errors.New("you're not allowed to access this data")))
			return
		}

		c.Next()
	}
}
