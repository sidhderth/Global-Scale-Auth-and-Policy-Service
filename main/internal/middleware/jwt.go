package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
)

type JWTMiddleware struct{
	keySet jwk.Set
}

func NewJWTMiddleware(jwksURL string) (*JWTMiddleware, error) {
	set, err := jwk.Fetch(context.Background(), jwksURL)
	if err != nil {
		return nil, err
	}
	return &JWTMiddleware{}, nil
}

func (m *JWTMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error" : "missing/malformed token"})
			return
		}
		tokenString := strings.TrimPrefix("Bearer ", auth)
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			kid := t.Header.Get("kid")
			key, ok = m.keySet.LookupKeyID(kid)
			if !ok {
				return nil, jwt.ErrTokenMalformed
			}
			var pubkey interface{}
			return pubkey, key.Raw(&pubkey)
		})
		if err != nil || !token.valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error" : "invalid token"})
			return
		}
		c.Set("claims", token.Claims.(jwt.MapClaims))
		c.Next()
	}
}
