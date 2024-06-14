package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)          // GET /events
	server.GET("/events/:id", getEvent)       // GET /events/:id
	server.POST("/events", createEvent)       // POST /events
	server.PUT("/events/:id", updateEvent)    // PUT /events/:id
	server.DELETE("/events/:id", deleteEvent) // DELETE /events/:id
	server.POST("/signup")                    // POST /signup
}
