package hatetepe

import "log"

type Router struct {
	handlers map[string]map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]map[string]Handler),
	}
}

func (r *Router) Handle(method, pattern string, handler Handler) {
	if _, ok := r.handlers[pattern]; !ok {
		r.handlers[pattern] = make(map[string]Handler)
	}
	r.handlers[pattern][method] = handler
}

func (r *Router) HandleFunc(method, pattern string, handler func(ResponseWriter, *Request)) {
	r.Handle(method, pattern, HandlerFunc(handler))
}

func (r *Router) ServeHTTP(w ResponseWriter, req *Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic in handler:", r)
			w.WriteHeader(500)
			w.Write([]byte("Internal Server Error"))
		}
	}()

	if handlers, ok := r.handlers[req.URL]; ok {
		if handler, ok := handlers[req.Method]; ok {
			handler.ServeHTTP(w, req)
			return
		}
	}

	// If no handler is found, return 404
	w.WriteHeader(404)
	w.Write([]byte("404 Not Found"))
}
