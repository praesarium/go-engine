package engine

import (
	"net/http"
	"sync"
)

type (

	Middleware func(*Context)

	engine struct {
		mware  []Middleware
		router *router
		pooler sync.Pool
	}
)

func (e *engine) build(h1 []Middleware, h2 []Middleware) []Middleware {

	handlers := make([]Middleware, 0, len(h1)+len(h2))

	handlers = append(handlers, h1...)
	handlers = append(handlers, h2...)

	return handlers
}

func CreateEngine() *engine {

	engine := &engine{}
	engine.router = createRouter(engine)

	engine.pooler.New = func() interface{} {
		return &Context{engine: engine}
	}

	return engine
}

func (e *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// push to driver handler
	e.router.driver.ServeHTTP(w, r)
}

func (e *engine) Use(mware ...Middleware) {

	// add middleware to global stack
	e.mware = append(e.mware, mware...)
}

func (e *engine) ListenTLS(h string, p string, c string, k string) error {

	addr := h + ":" + p

	if err := http.ListenAndServeTLS(addr, c, k, e); err != nil {
		return err
	}

	return nil
}

func (e *engine) Listen(h string, p string) error {

	addr := h + ":" + p

	if err := http.ListenAndServe(addr, e); err != nil {
		return err
	}

	return nil
}

func (e *engine) HandleRoutes(routes []Route) {

	// push to router
	e.router.handleRoutes(routes)
}
