package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = createRender()
	r.Static("/public", "./public")
	routes(r)
	log.Panic(r.Run(":8080"))
}
