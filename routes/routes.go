package routes

import (
	"github.com/JoaoFerrareis02/REST-API-GO/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)    // GET /events
	server.GET("/events/:id", getEvent) // GET /events/:id

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvent)       // POST /events
	authenticated.PUT("/events/:id", updateEvent)    // PUT /events/:id
	authenticated.DELETE("/events/:id", deleteEvent) // DELETE /events/:id

	server.POST("/signup", signup) // POST /signup
	server.POST("/login", login)   // POST /login
}
