package server

import (
	"actor-demo/internal/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//--------------------------------------//
// this file including handler for actor//
//--------------------------------------//

func (s *Server) createActorHandler(c *gin.Context) {
	var actorReq model.ActorRequest

	if err := c.ShouldBindJSON(&actorReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	actor, err := s.db.CreateActor(actorReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, actor)
}

func (s *Server) getAllActorHandler(c *gin.Context) {
	actors, err := s.db.GetAllActor()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if actors == nil {
		actors = []model.Actor{}
	}
	c.JSON(http.StatusOK, actors)
}

func (s *Server) getActorHandler(c *gin.Context) {
	rawId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required").Error()})
	}

	parsedId, err := strconv.Atoi(rawId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actor, err := s.db.GetActorById(int64(parsedId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actor)
}

func (s *Server) updateActorHandler(c *gin.Context) {
	rawId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required").Error()})
	}

	parsedId, err := strconv.Atoi(rawId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req model.ActorRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ID = int64(parsedId)

	actor, err := s.db.UpdateActor(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actor)
}

func (s *Server) deleteActorHandler(c *gin.Context) {
	rawId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required").Error()})
	}

	parsedId, err := strconv.Atoi(rawId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.db.DeleteActorById(int64(parsedId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
