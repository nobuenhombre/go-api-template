package server

import (
	"errors"
	"fmt"
	"go-api-template/src/internal/app/api-my-domain-com/api/server/config"
	"go-api-template/src/internal/app/api-my-domain-com/api/server/router"
	"log"
	"net/http"
	"os"
)

type HTTPServer struct {
	Router *router.HTTPRouter
	Server *http.Server
}

func NewHTTPServer(config *config.HTTPServerConfig, logFile *os.File) (srv *HTTPServer) {
	srv = new(HTTPServer)

	srv.Router = router.NewHTTPRouter(logFile, &config.Store)

	srv.Server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler: srv.Router.Router,
	}

	return
}

func (srv *HTTPServer) Run() {
	go func() {
		err := srv.Server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	srv.gracefulShutDown()
}
