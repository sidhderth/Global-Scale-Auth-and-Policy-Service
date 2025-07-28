package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/sidhderth/main/internal/handlers"
	"github.com/sidhderth/main/internal/middleware"
)

func main() {
	router := gin.Default()
	//TODO: add router.Use(jwtMw.Handler()), router.Use(opaMw)

	router.GET("/hello", handlers.HelloREST)

	port := os.Getnv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on : %s", port)
	router.Run("0.0.0.0" + port)
}