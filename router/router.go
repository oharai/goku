package router

import (
	"net/http"
)

type Router struct {
}

type Route struct {
	method  string
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
	before  func(w http.ResponseWriter, r *http.Request)
	after   func(w http.ResponseWriter, r *http.Request)
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Run(addr string) error {
	return http.ListenAndServe(addr, nil)
}

func (r *Router) GET(path string, handler func(w http.ResponseWriter, r *http.Request)) *Route {
	return r.handle("GET", path, handler)
}

func (r *Router) POST(path string, handler func(w http.ResponseWriter, r *http.Request)) *Route {
	return r.handle("POST", path, handler)
}

func (r *Router) PUT(path string, handler func(w http.ResponseWriter, r *http.Request)) *Route {
	return r.handle("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler func(w http.ResponseWriter, r *http.Request)) *Route {
	return r.handle("DELETE", path, handler)
}

func (r *Router) HEAD(path string, handler func(w http.ResponseWriter, r *http.Request)) *Route {
	return r.handle("HEAD", path, handler)
}

func (r *Router) OPTIONS(path string, handler func(w http.ResponseWriter, r *http.Request)) *Route {
	return r.handle("OPTIONS", path, handler)
}

func (r *Router) handle(method, path string, handler func(w http.ResponseWriter, r *http.Request)) *Route {
	route := &Route{
		method:  method,
		path:    path,
		handler: handler,
	}
	route.register()
	return route
}

func (r *Route) register() {
	http.HandleFunc(r.path, func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != r.method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if r.before != nil {
			r.before(writer, request)
		}

		r.handler(writer, request)

		if r.after != nil {
			r.after(writer, request)
		}
	})
}

func (r *Route) Before(middlewareFunc func(w http.ResponseWriter, r *http.Request)) *Route {
	r.before = middlewareFunc
	return r
}

func (r *Route) After(middlewareFunc func(w http.ResponseWriter, r *http.Request)) *Route {
	r.after = middlewareFunc
	return r
}
