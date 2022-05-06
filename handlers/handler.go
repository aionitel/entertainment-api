package handlers

import "github.com/gin-gonic/gin"

type Handler struct{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) *Handler {
	h := &Handler{}
	h.Init(c)
	return h
}
