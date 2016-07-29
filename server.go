package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func main() {
	var data []Person

	InsertData()
	data, err := SearchPerson(bson.M{"name": "Bun"}, 0, 10)
	log.Print(data)
	if len(err) > 0 {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    data,
		})
	})
	r.Run()
}
