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
	return &JWTMiddleware{keySet : set}, nil
}

func (m *JWTMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error" : "missing/malformed token"})
			return
		}
		tokenString := strings.TrimPrefix(auth, "Bearer")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			kidVal, ok := t.Header["kid"]
			if !ok {
    			return nil, jwt.ErrTokenMalformed
			}
			kid, ok := kidVal.(string)
			if !ok {
			    return nil, jwt.ErrTokenMalformed
			}
			// then lookup:
			key, found := m.keySet.LookupKeyID(kid)

			if !found {
				return nil, jwt.ErrTokenMalformed
			}
			var pubkey interface{}
			return pubkey, key.Raw(&pubkey)
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error" : "invalid token"})
			return
		}
		c.Set("claims", token.Claims.(jwt.MapClaims))
		c.Next()
	}
}
