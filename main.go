package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Numbers struct {
	Number1 float32 `json:"number1"`
	Number2 float32 `json:"number2"`
}

func main() {
	router := gin.Default()

	// CORS middleware ekleme
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // İzin verilen origin'leri belirtin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	// Basit bir GET endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/add", func(c *gin.Context) {
		var numbers Numbers
		if err := c.ShouldBindJSON(&numbers); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var sumnumber = numbers.Number1 + numbers.Number2
		c.JSON(http.StatusOK, gin.H{
			"sum": sumnumber,
		})
	})
	router.POST("/sub", func(c *gin.Context) {
		var numbers Numbers
		if err := c.ShouldBindJSON(&numbers); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var sumnumber = numbers.Number1 - numbers.Number2
		c.JSON(http.StatusOK, gin.H{
			"sum": sumnumber,
		})
	})
	router.POST("/multi", func(c *gin.Context) {
		var numbers Numbers
		if err := c.ShouldBindJSON(&numbers); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var sumnumber = numbers.Number1 * numbers.Number2
		c.JSON(http.StatusOK, gin.H{
			"sum": sumnumber,
		})
	})
	router.POST("/divide", func(c *gin.Context) {
		var numbers Numbers
		if err := c.ShouldBindJSON(&numbers); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var sumnumber = numbers.Number1 / numbers.Number2
		c.JSON(http.StatusOK, gin.H{
			"sum": sumnumber,
		})
	})

	// API'yi çalıştırma
	router.Run(":8080")
}
