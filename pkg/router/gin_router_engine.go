package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouterEngine struct {
	r *gin.Engine
}

func NewGinRouterEngine() *GinRouterEngine {
	return &GinRouterEngine{}
}

func (g *GinRouterEngine) getRouter() *gin.Engine {
	if g.r == nil {
		g.r = gin.Default()
	}

	return g.r
}

func (g *GinRouterEngine) responseWriter(c *gin.Context) http.ResponseWriter {
	return c.Writer
}

func (g *GinRouterEngine) request(c *gin.Context) *http.Request {
	return c.Request
}

func (g *GinRouterEngine) ginHandlerAdapter(handler http.HandlerFunc) func(*gin.Context) {
	return func(c *gin.Context) {
		handler(g.responseWriter(c), g.request(c))
	}
}

func (g *GinRouterEngine) Get(pattern string, handler http.HandlerFunc) {
	r := g.getRouter()
	r.GET(pattern, g.ginHandlerAdapter(handler))
}

func (g *GinRouterEngine) Post(pattern string, handler http.HandlerFunc) {
	r := g.getRouter()
	r.POST(pattern, g.ginHandlerAdapter(handler))
}

func (g *GinRouterEngine) Put(pattern string, handler http.HandlerFunc) {
	r := g.getRouter()
	r.PUT(pattern, g.ginHandlerAdapter(handler))
}

func (g *GinRouterEngine) Patch(pattern string, handler http.HandlerFunc) {
	r := g.getRouter()
	r.PATCH(pattern, g.ginHandlerAdapter(handler))
}

func (g *GinRouterEngine) Delete(pattern string, handler http.HandlerFunc) {
	r := g.getRouter()
	r.DELETE(pattern, g.ginHandlerAdapter(handler))
}

func (c *GinRouterEngine) Listen(port string) {
	r := c.getRouter()
	r.Run(port)
}
