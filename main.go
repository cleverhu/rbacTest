package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rbacTest/src/handlers"
)

func main() {
	r := gin.New()

	r.GET("/depts", handlers.Get)
	//r.POST("/", nil)
	//r.DELETE("/", nil)

	log.Fatal(r.Run(":80"))

}
