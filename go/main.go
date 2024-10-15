package main

import "github.com/gin-gonic/gin"

// main function
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello, World!",
		})
	})

	r.GET("/alert", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "It's fine",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
