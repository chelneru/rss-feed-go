package main

import (
	"github.com/gin-gonic/gin"
	"rss-feed/internal/services"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Second GET method
	r.GET("/rss-feed", func(c *gin.Context) {

		service := services.FeedService{}
		feeds, err := service.FetchFeeds("https://lorem-rss.herokuapp.com/feed?unit=day")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		c.JSON(200, feeds)

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
