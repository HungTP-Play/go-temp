package router

import "net/http"

type RouterConfigs struct {
	EngineType string
	Port       string
}

type Router struct {
	RouterConfigs
	engine RouterEngine
}

func NewRouter(engineType string, port string) *Router {
	return &Router{
		RouterConfigs: RouterConfigs{
			Port:       port,
			EngineType: engineType,
		},
	}
}

func (r *Router) createEngine() RouterEngine {
	defaultEngine := NewChiRouterEngine()
	if r.EngineType == "gin" {
		return NewGinRouterEngine()
	}

	return defaultEngine
}

func (r *Router) getEngine() RouterEngine {
	if r.engine == nil {
		r.engine = r.createEngine()
	}

	return r.engine
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	e := r.getEngine()
	e.Get(path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	e := r.getEngine()
	e.Post(path, handler)
}

func (r *Router) Patch(path string, handler http.HandlerFunc) {
	e := r.getEngine()
	e.Patch(path, handler)
}

func (r *Router) Delete(path string, handler http.HandlerFunc) {
	e := r.getEngine()
	e.Delete(path, handler)
}

func (r *Router) Listen() {
	e := r.getEngine()
	e.Listen(r.Port)
}
