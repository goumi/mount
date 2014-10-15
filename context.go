package mount

import "github.com/goumi/web"

// Extend context interface to contain a path method
type Context interface {
	web.Context

	// Get original request path
	Path() string
}

// context adds an extra path to the previous context
type context struct {
	web.Context

	// Request original path
	path string
}

// Create a new context
func newContext(ctx web.Context) Context {
	return &context{
		Context: ctx,
		path:    ctx.Request().URL.Path,
	}
}

// Next() restores the path and calls next middleware
func (ctx *context) Next() {

	// Set the path back to the original
	ctx.Request().URL.Path = ctx.path

	// Next
	ctx.Context.Next()
}

// Path() access the path
func (ctx *context) Path() string {
	return ctx.path
}
