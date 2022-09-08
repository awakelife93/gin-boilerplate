package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awakelife93/gin-boilerplate/src/lib"
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	"github.com/awakelife93/gin-boilerplate/src/lib/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnvironmentFile() {
	// * enter the env file name of your project.
	error := godotenv.Load(".env")

	if error != nil {
		log.Fatal(error.Error())
	}
}

// * pre loading application level env
func init() {
	loadEnvironmentFile()
}

func main() {
	// * ReleaseMode (release) or DebugMode (debug)
	// * default DebugMode (debug)
	var mode string = os.Getenv("env")

	if mode == "" {
		mode = gin.DebugMode
	}

	var engine *gin.Engine = lib.CreateEngine(mode)
	engine = routes.Initialize(engine)
	engine = middleware.Initialize(engine)

	fmt.Println("Start Gin Application")
	engine.Run(":8080")
}
