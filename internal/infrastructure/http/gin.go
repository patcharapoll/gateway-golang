package http

import (
	"bytes"
	"context"
	"fmt"
	"gateway-golang/internal/config"
	"gateway-golang/internal/graph/generated"
	"gateway-golang/internal/graph/resolver"
	"gateway-golang/internal/infrastructure/middleware"
	"gateway-golang/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHTTPServer)

// GinHTTPServer ...
type GinHTTPServer struct {
	config   *config.Configuration
	route    *gin.Engine
	resolver *resolver.Resolver
	server   *http.Server
	auth     *middleware.Auth
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write ...
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// NewHTTPServer ...
func NewHTTPServer(
	config *config.Configuration,
	resolver *resolver.Resolver,
) *GinHTTPServer {
	r := gin.New()
	r.Use(utils.GinContextToContextMiddleware(), middleware.HeaderMiddleware())
	//r.Use(func(c *gin.Context) {
	//	//var bodyBytes []byte
	//	//if c.Request.Body != nil {
	//	//	bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	//	//}
	//	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	//	c.Writer = blw
	//	c.Next()
	//	fmt.Println("user agent --> ", c.Request.UserAgent())
	//	fmt.Println("body --> ", blw.body.String())
	//	fmt.Println("referer --> ",c.Request.Referer())
	//})
	fmt.Println("referer --> ")
	r.Use(gin.Recovery())
	return &GinHTTPServer{
		route:    r,
		config:   config,
		resolver: resolver,
	}
}

func (g *GinHTTPServer) configure() *gin.Engine {
	r := g.route

	r.GET("/healthz", func(c *gin.Context) {
		c.String(200, "OK")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		srv := playground.Handler("GraphQL playground", "/query")
		srv.ServeHTTP(c.Writer, c.Request)
	})

	r.POST("/query", func(c *gin.Context) {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: g.resolver}))
		srv.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/query", func(c *gin.Context) {
		srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: g.resolver}))
		srv.AddTransport(transport.POST{})
		srv.AddTransport(transport.Websocket{
			KeepAlivePingInterval: 10 * time.Second,
			Upgrader: websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					// Check against your desired domains here
					return true
				},
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			},
		})
		srv.Use(extension.Introspection{})
		srv.ServeHTTP(c.Writer, c.Request)
	})

	return r
}

// Start ...
func (g *GinHTTPServer) Start(_ context.Context) {
	router := g.configure()
	g.server = &http.Server{
		Addr:    ":" + g.config.Port,
		Handler: router,
	}

	go func() {
		if err := g.server.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n\n", err)
		}
	}()

	log.Println("Listening and serving HTTP on", g.config.Port)
}

// Stop ...
func (g *GinHTTPServer) Stop(ctx context.Context) error {
	return g.server.Shutdown(ctx)
}
