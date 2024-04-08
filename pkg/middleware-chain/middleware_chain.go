package middleware_chain

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Constructor is type of httprouter handler
type Constructor func(httprouter.Handle) httprouter.Handle

// Chain is struck for list of middleware
type Chain struct {
	constructors []Constructor
}

// New is for innitial new chain of
func New(constructors ...Constructor) Chain {
	return Chain{append(([]Constructor)(nil), constructors...)}
}

// Then is for http router handler
func (c Chain) Then(h httprouter.Handle) httprouter.Handle {
	if h == nil {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {}
	}
	for i := range c.constructors {
		h = c.constructors[len(c.constructors)-1-i](h)
	}

	return h
}

// Append is for add chain router handler
func (c Chain) Append(constructors ...Constructor) Chain {
	newCons := make([]Constructor, 0, len(c.constructors)+len(constructors))
	newCons = append(newCons, c.constructors...)
	newCons = append(newCons, constructors...)

	return Chain{newCons}
}
