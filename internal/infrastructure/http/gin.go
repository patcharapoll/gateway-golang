package http

import (
	"bytes"
	"context"
	"gateway-golang/internal/config"
	"gateway-golang/internal/infrastructure/middleware"
	"gateway-golang/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewHTTPServer)

// GinHTTPServer ...
type GinHTTPServer struct {
	config *config.Configuration
	route  *gin.Engine
	server *http.Server
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
	r.Use(gin.Recovery())
	return &GinHTTPServer{
		route:  r,
		config: config,
	}
}

func (g *GinHTTPServer) configure() *gin.Engine {
	r := g.route

	r.GET("/healthz", func(c *gin.Context) {
		c.String(200, "OK")
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
