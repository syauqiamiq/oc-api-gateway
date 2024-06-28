package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"ocApiGateway/handler/courseHandler"
	"ocApiGateway/handler/mediaHandler"
	"ocApiGateway/handler/paymentHandler"
	"ocApiGateway/handler/userHandler"
	"ocApiGateway/middleware"
	"ocApiGateway/services/courseService"
	"ocApiGateway/services/mediaService"
	"ocApiGateway/services/paymentService"
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

	paymentService := paymentService.NewService()
	paymentHandler := paymentHandler.NewHandler(paymentService)

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

	// Course
	courseRouteV1 := routerV1.Group("/course")
	courseRouteV1.GET("/", courseHandler.GetCourseHandler)
	courseRouteV1.GET("/:id", courseHandler.GetCourseByIDHandler)
	courseRouteV1.POST("/", courseHandler.CreateCourseHandler)
	courseRouteV1.PUT("/:id", courseHandler.UpdateCourseByIDHandler)
	courseRouteV1.DELETE("/:id", courseHandler.DeleteCourseByIDHandler)

	// Chapter
	chapterRouteV1 := routerV1.Group("/chapter")
	chapterRouteV1.GET("/", courseHandler.GetChapterHandler)
	chapterRouteV1.GET("/:id", courseHandler.GetChapterByIDHandler)
	chapterRouteV1.POST("/", courseHandler.CreateChapterHandler)
	chapterRouteV1.PUT("/:id", courseHandler.UpdateChapterByIDHandler)
	chapterRouteV1.DELETE("/:id", courseHandler.DeleteChapterByIDHandler)

	// Lesson
	lessonRouteV1 := routerV1.Group("/lesson")
	lessonRouteV1.GET("/", courseHandler.GetLessonHandler)
	lessonRouteV1.GET("/:id", courseHandler.GetLessonByIDHandler)
	lessonRouteV1.POST("/", courseHandler.CreateLessonHandler)
	lessonRouteV1.PUT("/:id", courseHandler.UpdateLessonByIDHandler)
	lessonRouteV1.DELETE("/:id", courseHandler.DeleteLessonByIDHandler)

	// My Course
	myCourseRouteV1 := routerV1.Group("/my-course")
	myCourseRouteV1.GET("/", courseHandler.GetMyCourseHandler)

	// Order
	orderRouteV1 := routerV1.Group("/order")
	orderRouteV1.GET("/", paymentHandler.GetOrderHandler)
	orderRouteV1.POST("/", paymentHandler.CheckoutOrderHandler)

	err = router.Run(fmt.Sprintf(":%s", os.Getenv("RUNNING_PORT")))
	if err != nil {
		panic("Error When Running")
	}
}
