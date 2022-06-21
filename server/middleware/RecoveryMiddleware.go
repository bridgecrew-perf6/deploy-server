package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"oneclick/server/controller/methods"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				methods.ErrorDesc(c, fmt.Sprint(err))
			}
		}()
		c.Next()
	}
}
