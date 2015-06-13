package engine

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (

	Route struct {
		Middleware []Middleware
		Method   string
		Pattern  string
		Name	 string
	}

	router struct {
		driver *httprouter.Router
		engine *engine
	}
)

func createRouter(e *engine) *router {

	newRouter := &router{}
	newRouter.driver = httprouter.New()
	newRouter.engine = e

	newRouter.driver.MethodNotAllowed = func(

		writer http.ResponseWriter, request *http.Request) {
		ctx := createContext(e, writer, request, nil, e.mware, "no_method")

		ctx.Writer.WriteHeader(http.StatusMethodNotAllowed)
		ctx.NextMiddleware()
		ctx.recycle()
	}

	newRouter.driver.NotFound = func(

		writer http.ResponseWriter, request *http.Request) {
		ctx := createContext(e, writer, request, nil, e.mware, "no_route")

		ctx.Writer.WriteHeader(http.StatusNotFound)
		ctx.NextMiddleware()
		ctx.recycle()
	}

	return newRouter
}

func (r *router) handleRoute(route Route) {

	middleware := r.engine.build(
		r.engine.mware, route.Middleware,
	)

	r.driver.Handle(route.Method, route.Pattern,
		func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

			ctx := createContext(
				r.engine, writer, request, params, middleware, route.Name,
			)

			ctx.NextMiddleware()
			ctx.recycle()
		},
	)
}

func (r *router) handleRoutes(routes []Route) {

	for _, route := range routes {
		r.handleRoute(route)
	}
}
