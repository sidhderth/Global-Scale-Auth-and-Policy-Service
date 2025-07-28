package middleware

import "github.com/gin-gonic/gin"

func OPAMiddleware(policyPath string) gin.HanlderFunc {
	return func(c *gin.Context) {c.Next()}
}