package route

import (
	"net/http"
)

// Builder build http.ServeMux to map URL(Pattern) to http.Handler
type Builder struct {
	route map[string]http.Handler
}

// MapRoute route pattern to handler
func (b *Builder) MapRoute(pattern string, handler http.Handler) *Builder {
	b.route[pattern] = handler

	return b
}

// MapRouteFunc route pattern to handler
func (b *Builder) MapRouteFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) *Builder {
	b.route[pattern] = http.HandlerFunc(handler)

	return b
}

// BuildServeMux buil http.ServeMux instance
func (b *Builder) BuildServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	for pattern, handler := range b.route {
		mux.Handle(pattern, handler)
	}

	return mux
}

var builder *Builder

// MapRoute route pattern to handler
func MapRoute(pattern string, handler http.Handler) *Builder {
	return builder.MapRoute(pattern, handler)
}

// MapRouteFunc route pattern to handler
func MapRouteFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) *Builder {
	return builder.MapRouteFunc(pattern, handler)
}

// BuildServeMux build http.ServeMux instance
func BuildServeMux() *http.ServeMux {
	return builder.BuildServeMux()
}

func init() {
	builder = &Builder{route: make(map[string]http.Handler)}
	builder.MapRouteFunc("/", http.NotFound)
}
