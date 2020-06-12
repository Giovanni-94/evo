package evo

import (
	"github.com/gofiber/fiber"
)

type group struct {
	app *fiber.Group
}

// Group is used for Routes with common prefix to define a new sub-router with optional middleware.
func Group(path string, handlers ...func(request *Request)) *group {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Group
	if len(handlers) > 0 {
		for _, handler := range handlers {
			route = app.Group(path, func(ctx *fiber.Ctx) {
				r := Upgrade(ctx)
				handler(r)
			})
		}
	} else {
		route = app.Group(path)
	}
	gp := group{
		app: route,
	}
	return &gp
}

/*func Group(prefix string, handlers ...func(*fiber.Ctx)) *fiber.Group {
	if app == nil {
		panic("Access object before call Setup()")
	}
	return app.Group(prefix, handlers...)
}*/

// Static append path with given prefix to static files
func Static(prefix, path string) {
	statics = append(statics, [2]string{prefix, path})
}

// Use registers a middleware route.
// Middleware matches requests beginning with the provided prefix.
// Providing a prefix is optional, it defaults to "/"
func Use(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Use(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// Connect : https://fiber.wiki/application#http-methods
func Connect(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Connect(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// Put : https://fiber.wiki/application#http-methods
func Put(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Put(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// Post : https://fiber.wiki/application#http-methods
func Post(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Post(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// Delete : https://fiber.wiki/application#http-methods
func Delete(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Delete(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// Head : https://fiber.wiki/application#http-methods
func Head(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Head(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// Patch : https://fiber.wiki/application#http-methods
func Patch(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Patch(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// Options : https://fiber.wiki/application#http-methods
func Options(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Options(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// Trace : https://fiber.wiki/application#http-methods
func Trace(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Trace(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// Get : https://fiber.wiki/application#http-methods
func Get(path string, handlers ...func(request *Request)) *fiber.Route {
	if app == nil {
		panic("Access object before call Setup()")
	}
	var route *fiber.Route
	for _, handler := range handlers {
		route = app.Get(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}
	return route
}

// All : https://fiber.wiki/application#http-methods
func All(path string, handlers ...func(request *Request)) {
	if app == nil {
		panic("Access object before call Setup()")
	}

	for _, handler := range handlers {
		app.All(path, func(ctx *fiber.Ctx) {
			r := Upgrade(ctx)
			handler(r)
		})
	}

}

// Shutdown gracefully shuts down the server without interrupting any active connections.
// Shutdown works by first closing all open listeners and then waiting indefinitely for all connections to return to idle and then shut down.
//
// When Shutdown is called, Serve, ListenAndServe, and ListenAndServeTLS immediately return nil.
// Make sure the program doesn't exit and waits instead for Shutdown to return.
//
// Shutdown does not close keepalive connections so its recommended to set ReadTimeout to something else than 0.
func Shutdown() error {
	if app == nil {
		panic("Access object before call Setup()")
	}
	return app.Shutdown()
}
