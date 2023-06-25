package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Redis address
		Password: "",           // Redis password, leave empty if no authentication is required
		DB:       0,            // Redis database number
	})

	// Set a value in Redis
	err := rdb.Set(ctx, "mykey", "Hello Redis", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// Get the value from Redis
	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Value from Redis:", val)

	// Initialize Gin router
	router := gin.Default()

	// Define a /ping route
	router.GET("/ping", func(c *gin.Context) {
		// Get the value from Redis
		val, err := rdb.Get(ctx, "mykey").Result()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, val)
	})

	// Run the server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
