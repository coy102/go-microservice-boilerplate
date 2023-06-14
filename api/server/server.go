package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-microservices.org/api/graph"
	"go-microservices.org/api/graph/resolver"
	"go-microservices.org/api/login"
	loginresolver "go-microservices.org/api/login/resolver"
	"go-microservices.org/api/middleware"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

// Server as API server
type Server struct {
	instance *http.Server
	port     string
	useTLS   bool
	tlsCert  string
	tlsKey   string
	logger   *log.Logger
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found ..")
	}

}

func main() {
	srv := createServer()
	srv.writeLog("Server is starting...")

	ctx := srv.handleShutdown()
	srv.start()

	<-ctx.Done()
	srv.writeLog("Server stopped...")
}

func createServer() *Server {
	srv := new(Server)

	srv.port = defaultPort
	if port := os.Getenv("PORT"); port != "" {
		srv.port = port
	}

	srv.logger = log.New(os.Stdout, "[http] ", log.LstdFlags)
	srv.instance = &http.Server{
		Addr:     ":" + srv.port,
		Handler:  buildRouter(),
		ErrorLog: srv.logger,
	}

	return srv
}

func buildRouter() http.Handler {
	router := chi.NewRouter()
	router.Use(
		middleware.GetCorsAPI().Handler,
		middleware.AuthCheck,
	)

	options := []handler.Option{
		handler.WebsocketUpgrader(
			websocket.Upgrader{
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			}),
		handler.CacheSize(0),
	}

	router.Handle(
		"/",
		handler.Playground("GraphQL playground", "/login"),
	)

	router.Handle(
		"/login",
		handler.GraphQL(
			login.NewExecutableSchema(
				login.Config{
					Resolvers: &loginresolver.Resolver{},
				},
			),
			options...,
		),
	)

	router.Handle(
		"/query",
		handler.GraphQL(
			graph.NewExecutableSchema(
				graph.Config{
					Resolvers: &resolver.Resolver{},
					Directives: graph.DirectiveRoot{
						HasPermission: middleware.PermissionCheck,
					},
				},
			),
			options...,
		),
	)

	return router
}

func (s *Server) start() {
	s.writeLog("Server is ready to handle requests at %q", s.instance.Addr)
	s.writeLog("Connect to http://localhost:%s/ for GraphQL playground", s.port)

	var err error
	if s.useTLS {
		err = s.instance.ListenAndServeTLS(s.tlsCert, s.tlsKey)
	} else {
		err = s.instance.ListenAndServe()
	}

	if err != http.ErrServerClosed {
		s.fatal("Could not listen on %q: %v", s.instance.Addr, err)
	}
}

func (s *Server) handleShutdown() context.Context {
	ctx, done := context.WithCancel(context.Background())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer done()

		<-quit
		signal.Stop(quit)
		close(quit)

		s.writeLog("Server is shutting down...")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		s.instance.SetKeepAlivesEnabled(false)
		if err := s.instance.Shutdown(ctx); err != nil {
			s.fatal("Could not gracefully shutdown the server: %v", err)
		}
	}()

	return ctx
}

func (s *Server) fatal(msg string, args ...interface{}) {
	s.logger.Fatalf(msg, args...)
}

func (s *Server) writeLog(msg string, args ...interface{}) {
	s.logger.Printf(msg, args...)
}
