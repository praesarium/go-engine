package engine

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Context struct {
	Route		string
	Params  httprouter.Params
	Request *http.Request
	Writer  writerInterface

	mware		[]Middleware
	engine	*engine
	store		map[string]interface{}
	index		int8
	abort		bool
}

func createContext(

	engine 	*engine,
	writer 	http.ResponseWriter,
	request *http.Request,
	params 	httprouter.Params,
	mware 	[]Middleware,
	route 	string,

) *Context {

	// get context from engine pool
	ctx := engine.pooler.Get().(*Context)

	ctx.store = nil
	ctx.mware = mware
	ctx.index = -1

	ctx.Route   = route
	ctx.Writer  = createWriter(writer)
	ctx.Request = request
	ctx.Params  = params

	return ctx
}

func (c *Context) recycle() {
	c.engine.pooler.Put(c)
}

func (c *Context) Render(status int, driver Renderer, data interface{}) error {
	return driver.Render(c.Writer, status, data)
}

func (c *Context) GetValue(k string) interface{} {

	if c.store != nil {

		value, ok := c.store[k]
		if ok {
			return value
		}
	}

	return nil
}

func (c *Context) SetValue(key string, val interface{}) {

	if c.store == nil {
		c.store = make(map[string]interface{})
	}

	c.store[key] = val
}

func (c *Context) NextMiddleware() {

	count := int8(len(c.mware))
	c.index += 1

	if c.abort != true && c.index < count {
		c.mware[c.index](c)
	}
}

func (c *Context) Parse(driver Parser, data interface{}) error {
	return driver.Parse(c.Request, data)
}

func (c *Context) Abort(status int) {

	c.Writer.WriteHeader(status)
	c.abort = true
}
