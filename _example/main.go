package main

import (
	"github.com/gin-gonic/gin"

	"github.com/wyy-go/wnocache"
)

func main() {
	g := gin.Default()
	g.Use(wnocache.New())
	g.Run(":8080")
}
