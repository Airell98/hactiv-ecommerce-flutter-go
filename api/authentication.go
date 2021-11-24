package api

import (
	"errors"
	"net/http"

	"github.com/AirellJordan98/hacktiv_ecommerce/util"
	"github.com/gin-gonic/gin"
)

func authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := util.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse("Unathenticated", errors.New("invalid token")))
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
