package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awakelife93/gin-boilerplate/src/config"
	"github.com/awakelife93/gin-boilerplate/src/lib/database"
	"github.com/awakelife93/gin-boilerplate/src/lib/middleware"
	"github.com/awakelife93/gin-boilerplate/src/lib/routes"
	"github.com/awakelife93/gin-boilerplate/src/repositories"
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

func setModeHandler() {
	// * ReleaseMode (release) or DebugMode (debug)
	// * default DebugMode (debug)
	mode := os.Getenv("env")

	if mode == "" {
		mode = gin.DebugMode
	}

	if mode == gin.DebugMode {
		gin.ForceConsoleColor()
	} else {
		gin.DisableConsoleColor()
	}

	gin.SetMode(mode)
}

func init() {
	config.Initialize()
	loadEnvironmentFile()
	setModeHandler()
}

func main() {
	engine := gin.Default()

	// * initialize database
	database.Initialize()

	// * initialize database
	database.GenerateSampleData()

	// * initialize repositories
	repositories.Initialize()

	// * initialize route
	routes.Initialize(engine)

	// * initialize middleware
	middleware.Initialize(engine)

	fmt.Println("Start Gin Application")
	engine.Run(":8080")
}
