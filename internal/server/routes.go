package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	actor := r.Group("/actors")
	{
		actor.POST("", s.createActorHandler)
		actor.GET("", s.getAllActorHandler)
		actor.GET("/:id", s.getActorHandler)
		actor.PUT("/:id", s.updateActorHandler)
		actor.DELETE("/:id", s.deleteActorHandler)
	}
	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
