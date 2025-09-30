package handlers

import "github.com/gin-gonic/gin"

func HelloREST(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, world!"})
}