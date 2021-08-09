package main

import "net/http"

type handlerFunc func(ctx *context)

type context struct {
	response http.ResponseWriter
	request  *http.Request
}

type Engine struct {
	router map[string]handlerFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]handlerFunc),
	}
}

func (p *Engine) Get(path string, handler handlerFunc) {
	if _, ok := p.router["GET"+path]; ok {
		return
	}

	p.router["GET"+path] = handler
}

func (p *Engine) POST(path string, handler handlerFunc) {
	if _, ok := p.router["POST"+path]; ok {
		return
	}

	p.router["POST"+path] = handler
}

func (p *Engine) Run(addr string) {
	http.ListenAndServe(addr, p)
}

func (p *Engine) ServeHTTP(r http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	if handler, ok := p.router[method+path]; ok {
		handler(r, req)
	}
}
