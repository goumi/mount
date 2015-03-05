package mount

import "github.com/goumi/web"

// handler extends web handler to have a pattern
type handler struct {
	web.Handler

	// Pattern to match
	pattern string
}

// New mounting point
func New(p string, h web.Handler) web.Handler {
	return &handler{
		Handler: h,
		pattern: p,
	}
}

// Serve creates a new context in order to remap the URL Path
func (h *handler) Serve(ctx web.Context) {

	// Load the matched string
	p := match(h.pattern, ctx.Request().URL.Path)

	// Path doesn't match
	if p == "" {

		// We should just continue the chain
		ctx.Next()

		// Skip the rest
		return
	}

	// Create new context
	ctx = NewContext(ctx, p)

	// Serve the hadler
	h.Handler.Serve(ctx)
}

// match checks if there is a match on the pattern and on the request
func match(pattern, path string) (p string) {

	// Response is nil by default
	p = ""

	// Load the pattern length
	n := len(pattern)

	// No empty patterns
	if n == 0 {
		return
	}

	// Patterns should have a trailing slash in order to match anyting
	if pattern[n-1] != '/' {

		// Pattern is different
		if pattern == path {
			p = "/"
		}

		return
	}

	// Check if the pattern matches
	if len(path) >= n && path[0:n] == pattern {
		p = path[n-1:]
	}

	return
}
