// main.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	bind := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"result": "success"})
	})

	router.Run(bind)
}
