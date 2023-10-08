package handlers

import (
	"go-api-template/src/internal/app/api-my-domain-com/api/server/config"
	"go-api-template/src/internal/app/api-my-domain-com/api/server/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	Store *store.Store
}

func NewHttpHandler(redisConfig *config.RedisConfig) (handler *HttpHandler) {
	handler = new(HttpHandler)
	handler.Store = store.NewStore(redisConfig)
	return handler
}

func (h *HttpHandler) DefaultHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome API Server")
}

/**
 * Add New Handlers Here
 */
