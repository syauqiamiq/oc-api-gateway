package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"ocApiGateway/handler/courseHandler"
	"ocApiGateway/handler/mediaHandler"
	"ocApiGateway/handler/userHandler"
	"ocApiGateway/middleware"
	"ocApiGateway/services/courseService"
	"ocApiGateway/services/mediaService"
	"ocApiGateway/services/userService"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getCookieStore() []byte {
	h := sha1.New()
	return h.Sum([]byte(time.Now().String()))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mediaService := mediaService.NewService()
	mediaHandler := mediaHandler.NewHandler(mediaService)
	userService := userService.NewService()
	userHandler := userHandler.NewHandler(userService)

	courseService := courseService.NewService()
	courseHandler := courseHandler.NewHandler(courseService)

	router := gin.Default()

	router.Use(sessions.Sessions("SESSION-DATA", cookie.NewStore(getCookieStore())))

	routerV1 := router.Group("/api/v1")
	routerV1.POST("/register", userHandler.RegisterHandler)
	routerV1.POST("/login", userHandler.LoginHandler)
	routerV1.POST("/refresh-token", userHandler.RefreshTokenHandler)
	routerV1.POST("/logout", middleware.AuthMiddleware(), userHandler.LogoutHandler)
	routerV1.PUT("/my-profile", middleware.AuthMiddleware(), userHandler.UpdateUserHandler)
	routerV1.GET("/my-profile", middleware.AuthMiddleware(), userHandler.GetProfileHandler)

	// Media
	mediaRouteV1 := routerV1.Group("/media")
	mediaRouteV1.GET("/", mediaHandler.GetAllMediaHandler)
	mediaRouteV1.DELETE("/:id", mediaHandler.DeleteMediaByIdHandler)
	mediaRouteV1.POST("/", mediaHandler.UploadImageHandler)

	// Mentor
	mentorRouteV1 := routerV1.Group("/mentor")
	mentorRouteV1.GET("/", courseHandler.GetMentorHandler)
	mentorRouteV1.GET("/:id", courseHandler.GetMentorByIDHandler)
	mentorRouteV1.POST("/", courseHandler.CreateMentorHandler)
	mentorRouteV1.PUT("/:id", courseHandler.UpdateMentorByIDHandler)
	mentorRouteV1.DELETE("/:id", courseHandler.DeleteMentorByIDHandler)

	err = router.Run(fmt.Sprintf(":%s", os.Getenv("RUNNING_PORT")))
	if err != nil {
		panic("Error When Running")
	}
}
