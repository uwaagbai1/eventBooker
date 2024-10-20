package routes

import (
	"eventBooker/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Events Routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	// // Protected Events Routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	// User Routes
	server.POST("/signup", signUp)
	server.POST("/login", login)
	// Registration Routes
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/delete", cancelRegistration)
}
