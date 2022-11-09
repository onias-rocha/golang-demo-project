package controller

import (
	"github.com/gin-gonic/gin"
	"golang-softplayer-apply/service"
	"net/http"
)

type Controller struct {
	service service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) Get(g *gin.Context) {
	g.JSON(http.StatusOK, "pong")
}

func (c *Controller) GetAllPeople(g *gin.Context) {
	g.JSON(http.StatusOK, c.service.GetAllPeople())
}
