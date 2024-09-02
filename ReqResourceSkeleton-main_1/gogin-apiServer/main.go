package main

import (
	"log"
	"os"
	"sync"

	f "main/function"
	"main/handler"

	"github.com/gin-gonic/gin"
)

const JSON_SERVICE_PORT = ":9000"
const YAML_SERVICE_PORT = ":9001"

var wg sync.WaitGroup

func init() {
	os.Setenv("TZ", "Asia/Seoul")

	wg.Add(3)
}

// JSON Middleware is to set all types of request are JSON.
func JSONMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	}
}

// YAML Middleware is to set all types of request are YAML.
func YAMLMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/x-yaml")
		ctx.Next()
	}
}

func json_newServer() *gin.Engine {
	router := gin.Default()
	router.Use(JSONMiddleware())

	// general handler
	router.POST("/test", handler.EchoTestHandler)
	return router
}

func yaml_newServer() *gin.Engine {
	router := gin.Default()
	router.Use(YAMLMiddleware())

	// general handler
	router.POST("/resource", handler.ResourceHandler)
	router.POST("/final", handler.FinalHandler)

	return router
}

func main() {

	log.Println(f.GetTimeAndFormat("second"))

	// json_newServer().Run(JSON_SERVICE_PORT)
	yaml_newServer().Run(YAML_SERVICE_PORT)

	wg.Wait()
}
