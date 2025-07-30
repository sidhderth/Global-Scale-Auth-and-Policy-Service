package middleware

import "github.com/gin-gonic/gin"

//Placeholder
type JWTMiddleware struct{}

func NewJWTMiddleware(jwksURL string) (*JWTMiddleware, error) {
	return &JWTMiddleware{}, nil
}

func (m *JWTMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {c.Next()}
}
