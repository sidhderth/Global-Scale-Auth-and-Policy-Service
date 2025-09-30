package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/sidhderth/internal/handlers"
	"github.com/sidhderth/internal/middleware"
)

func main() {
	jwks := os.Getenv("JWKS_URL")
	if jwks == "" {
		jwks = "http://localhost:8081/realms/demo/protocol/openid-connect/certs"
	}
	jwtMw, err := middleware.NewJWTMiddleware(jwks)
	if err != nil {
		log.Fatalf("JWT Middleware init error : %v", err)
	}

	opaMw, err := middleware.NewOPAMiddleware("policy/authz.rego")
	if err != nil {
		log.Fatalf("OPA Middleware init error: %v", err)
	}

	router := gin.Default()
	router.Use(jwtMw.Handler())
	router.Use(opaMw.Handler())

	router.GET("/hello", handlers.HelloREST)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on : %s", port)
	router.Run("0.0.0.0" + port)
}