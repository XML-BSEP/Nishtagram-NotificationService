package router

import (
	"github.com/gin-gonic/gin"
	"notification-service/infrastructure/http/middleware"
	"notification-service/interactor"
)

func NewRoute(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.GET("/notification/:userId", handler.GetNotificationsByUserId)
	router.PUT("/notification/:notificationId", handler.UpdateNotificationStatus)
	router.GET("/proba", handler.GrpcProba)
	return router
}
