package main

import (
	"fmt"

	"github.com/awakelife93/gin-boilerplate/src/lib"
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	"github.com/awakelife93/gin-boilerplate/src/lib/routes"
	"github.com/gin-gonic/gin"
)

// todo: connect neo4j
func main() {

	fmt.Println("create engine")
	var engine *gin.Engine = lib.CreateEngine()

	fmt.Println("initialize router")
	engine = routes.Initialize(engine)

	engine = middleware.Initialize(engine)

	fmt.Println("Start Gin Application")
	engine.Run(":8080")
}
