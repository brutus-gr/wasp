package wasp

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
)

type Route struct {
	MessageName string
	Callback    Handler
}

type Routes = []Route
type Handler = func(m []byte, c net.Conn) error
type Middleware = func(m []byte, c net.Conn) ([]byte, error)

type Router struct {
	routes     Routes
	middleware []Middleware
}

// Handle handles the incoming TCP message
func (r *Router) Handle(m []byte, c net.Conn) error {
	// decode the message
	pm := Message{}
	e := json.NewDecoder(bytes.NewReader(m)).Decode(&pm)

	if e != nil {
		return e
	}

	// apply middleware
	for _, mw := range r.middleware {
		m, e = mw(m, c)
		if e != nil {
			return e
		}
	}

	// route it
	for _, r := range r.routes {
		if r.MessageName == pm.Name {
			return r.Callback(m, c)
		}
	}

	return errors.New("route not found")
}

func MakeRouter(routes Routes, middleware []Middleware) Router {
	return Router{
		routes:     routes,
		middleware: middleware,
	}
}
