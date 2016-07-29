package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	data := InsertData()
	out, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    out,
		})
	})
	r.Run()
}
