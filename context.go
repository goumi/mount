package mount

import "github.com/goumi/web"

// context adds an extra path to the previous context
type context struct {
	web.Context

	// Request original path
	origin string
}

// Create a new context
func NewContext(ctx web.Context, path string) web.Context {

	// Load the previous path
	origin := ctx.Request().URL.Path

	// Update the context path
	ctx.Request().URL.Path = path

	return &context{
		Context: ctx,
		origin:  origin,
	}
}

// Next() restores the path and calls next middleware
func (ctx *context) Next() {

	// Set the path back to the original
	ctx.Request().URL.Path = ctx.origin

	// Next
	ctx.Context.Next()
}
