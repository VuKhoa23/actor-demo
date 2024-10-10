package server

import (
	"actor-demo/internal/model"
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
		c.JSON(http.StatusBadRequest, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}

	actor, err := s.db.CreateActor(actorReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.HttpResponse[model.Actor]{Data: &actor, Success: true})
}

func (s *Server) getAllActorHandler(c *gin.Context) {
	actors, err := s.db.GetAllActor()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}
	if len(actors) == 0 {
		actors = []model.Actor{}
	}
	c.JSON(http.StatusOK, model.HttpResponse[[]model.Actor]{Success: true, Data: &actors})
}

func (s *Server) getActorHandler(c *gin.Context) {
	rawId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: "ID is required"})
		return
	}

	parsedId, err := strconv.Atoi(rawId)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}

	actor, err := s.db.GetActorById(int64(parsedId))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusOK, model.HttpResponse[model.Actor]{Data: nil, Success: true})
			return
		}
		c.JSON(http.StatusBadRequest, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[model.Actor]{Success: true, Data: &actor})
}

func (s *Server) updateActorHandler(c *gin.Context) {
	rawId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: "ID is required"})
		return
	}

	parsedId, err := strconv.Atoi(rawId)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}

	var req model.ActorRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}

	req.ID = int64(parsedId)

	actor, err := s.db.UpdateActor(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[model.Actor]{Success: true, Data: &actor})
}

func (s *Server) deleteActorHandler(c *gin.Context) {
	rawId, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: "ID is required"})
		return
	}

	parsedId, err := strconv.Atoi(rawId)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}

	err = s.db.DeleteActorById(int64(parsedId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.HttpResponse[model.Actor]{Data: nil, Success: false, Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.HttpResponse[model.Actor]{Data: nil, Success: true})
}
