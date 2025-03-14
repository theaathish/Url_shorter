package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"context"
)

var ctx = context.Background()

const URL_PREFIX = "http://localhost:8080/"
const SHORT_URL_LENGTH = 6

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	short := make([]rune, SHORT_URL_LENGTH)
	for i := range short {
		short[i] = letters[rand.Intn(len(letters))]
	}
	return string(short)
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	r := gin.Default()

	r.POST("/shorten", func(c *gin.Context) {
		var request struct {
			OriginalURL string `json:"original_url"`
			CustomAlias string `json:"custom_alias"`
			Expiry      int    `json:"expiry"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		shortURL := request.CustomAlias
		if shortURL == "" {
			shortURL = generateShortURL()
		}

		exists, _ := rdb.Exists(ctx, shortURL).Result()
		if exists > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "Custom alias already taken"})
			return
		}

		expiry := time.Duration(request.Expiry) * time.Second
		rdb.Set(ctx, shortURL, request.OriginalURL, expiry)

		c.JSON(http.StatusOK, gin.H{
			"short_url": URL_PREFIX + shortURL,
		})
	})

	r.GET("/:shortURL", func(c *gin.Context) {
		shortURL := c.Param("shortURL")
		originalURL, err := rdb.Get(ctx, shortURL).Result()
		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, originalURL)
	})

	r.Run(":8080")
}