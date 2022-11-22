package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"gihub.com/olamilekan000/bp-go/server/routes"
)

func Start() {
	envChecks()

	router := gin.New()

	routes.AddRoutes(router)

	port := os.Getenv("PORT")

	router.Run(port)
}

func envChecks() {
	fmt.Println("Checking env")
	envProps := []string{
		"PORT",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			log.Fatalf("Environment variable %s value is missing", k)
		}
	}
}
