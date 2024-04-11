package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ChiRouterEngine struct {
	router *chi.Mux
}

func NewChiRouterEngine() *ChiRouterEngine {
	return &ChiRouterEngine{}
}

func (c *ChiRouterEngine) getRouter() *chi.Mux {
	if c.router == nil {
		c.router = chi.NewRouter()
	}

	return c.router
}

func (c *ChiRouterEngine) Get(pattern string, handler http.HandlerFunc) {
	r := c.getRouter()
	r.Get(pattern, handler)
}

func (c *ChiRouterEngine) Post(pattern string, handler http.HandlerFunc) {
	r := c.getRouter()
	r.Post(pattern, handler)
}

func (c *ChiRouterEngine) Put(pattern string, handler http.HandlerFunc) {
	r := c.getRouter()
	r.Put(pattern, handler)
}

func (c *ChiRouterEngine) Patch(pattern string, handler http.HandlerFunc) {
	r := c.getRouter()
	r.Patch(pattern, handler)
}

func (c *ChiRouterEngine) Delete(pattern string, handler http.HandlerFunc) {
	r := c.getRouter()
	r.Delete(pattern, handler)
}

func (c *ChiRouterEngine) Listen(port string) {
	r := c.getRouter()
	http.ListenAndServe(port, r)
}
