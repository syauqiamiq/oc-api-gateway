package main

import (
	"fmt"
	"log"
	"ocApiGateway/handler/mediaHandler"
	"ocApiGateway/services/mediaService"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mediaService := mediaService.NewService()
	mediaHandler := mediaHandler.NewHandler(mediaService)

	router := gin.Default()

	routerV1 := router.Group("/api/v1")

	mediaRouteV1 := routerV1.Group("/media")

	mediaRouteV1.GET("/", mediaHandler.GetAllMediaHandler)
	mediaRouteV1.DELETE("/:id", mediaHandler.DeleteMediaByIdHandler)
	mediaRouteV1.POST("/", mediaHandler.UploadImageHandler)

	err = router.Run(fmt.Sprintf(":%s", os.Getenv("RUNNING_PORT")))
	if err != nil {
		panic("Error When Running")
	}
}
