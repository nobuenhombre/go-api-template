package router

import (
	"go-api-template/src/internal/app/api-my-domain-com/api/server/config"
	"go-api-template/src/internal/app/api-my-domain-com/api/server/router/handlers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

type HTTPRouter struct {
	Router   *gin.Engine
	Handlers *handlers.HttpHandler
}

func NewHTTPRouter(logFile *os.File, redisConfig *config.RedisConfig) (router *HTTPRouter) {
	router = new(HTTPRouter)
	router.WriteToLog(logFile)

	router.Handlers = handlers.NewHttpHandler(redisConfig)
	router.Router = gin.Default()
	router.AddRoutes()

	return
}

func (r *HTTPRouter) WriteToLog(logFile *os.File) {
	if logFile != nil {
		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(logFile)
	}
}

func (r *HTTPRouter) AddRoutes() {
	r.Router.GET("/", r.Handlers.DefaultHandler)

	/**
	 * Add New Routes Here
	 */
}
