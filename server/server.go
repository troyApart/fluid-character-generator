package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/troyApart/fluid-character-generator/endpoints"
	"github.com/urfave/negroni"
)

type Server struct {
	router     *mux.Router
	httpServer *http.Server
	middleware *negroni.Negroni

	host string
	port int
}

func (s *Server) initializeRoutes() {
	generateCharacterEndpoint := endpoints.NewGenerateCharacter()

	s.router.Handle("/v1/generate", negroni.New(
		negroni.Wrap(http.HandlerFunc(generateCharacterEndpoint.Get)),
	)).Methods("GET")
	s.router.Handle("/v1/generate", negroni.New(
		negroni.Wrap(http.HandlerFunc(generateCharacterEndpoint.Post)),
	)).Methods("POST")
	s.router.Handle("/v1/generate", negroni.New(
		negroni.Wrap(http.HandlerFunc(generateCharacterEndpoint.Patch)),
	)).Methods("PATCH")

	s.router.NotFoundHandler = http.HandlerFunc(endpoints.NotFound)
}

func (s *Server) initializeMiddleware() error {
	s.middleware.UseHandler(s.router)

	return nil
}

func New() (*Server, error) {
	s := &Server{
		router:     mux.NewRouter(),
		middleware: negroni.New(),
		host:       "localhost",
		port:       50000,
	}

	s.initializeRoutes()
	err := s.initializeMiddleware()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("error initializing middleware")
		return nil, err
	}
	return s, nil
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.httpServer = &http.Server{Addr: addr, Handler: s.middleware}

	log.Printf("rts listening on %s", addr)
	if err := s.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("error occurred when starting up rts %s", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
