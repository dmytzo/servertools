package mux

import (
	"fmt"
	"net/http"
)

type ServeMux struct {
	*http.ServeMux
}

func NewServeMux(routes ...RouteHandlers) *ServeMux {
	mux := http.NewServeMux()
	for _, route := range Routes(routes...) {
		mux.HandleFunc(fmt.Sprintf("%s %s", route.method, route.pattern), route.handler)
	}

	return &ServeMux{ServeMux: mux}
}

type RouteHandler struct {
	method  string
	pattern string
	handler http.HandlerFunc
}

type RouteHandlers []RouteHandler

func (r RouteHandlers) WithMiddlewares(middlewares ...func(http.HandlerFunc) http.HandlerFunc) RouteHandlers {
	for idx := range r {
		for _, middleware := range middlewares {
			r[idx].handler = middleware(r[idx].handler)
		}
	}

	return r
}

func Route(method string, pattern string, handler http.HandlerFunc) RouteHandlers {
	return RouteHandlers{{method: method, pattern: pattern, handler: handler}}
}

func Routes(routeGroups ...RouteHandlers) RouteHandlers {
	var flatRoutes RouteHandlers

	for _, routes := range routeGroups {
		for _, route := range routes {
			flatRoutes = append(flatRoutes, route)
		}
	}

	return flatRoutes
}
