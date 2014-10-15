package mount

import "github.com/goumi/web"

// Handler extends web handler to have a pattern
type handler struct {
	web.Handler

	// Pattern to match
	pattern string
}

func New(p string, h web.Handler) web.Handler {
	return &handler{
		Handler: h,
		pattern: p,
	}
}

// Serve() creates a new context in order to remap the URL Path
func (h *handler) Serve(ctx web.Context) {

	// Create new context
	mctx := newContext(ctx)

	// Load the matched string
	mp := match(h.pattern, mctx.Path())

	// Check if the path match
	if mp != "" {

		// Modify the request
		mctx.Request().URL.Path = mp

		// Serve the hadler
		h.Handler.Serve(mctx)
	}

	// Continue to the next one
	mctx.Next()
}

// match() checks if there is a match on the patter and on the request
func match(pattern, path string) (mp string) {

	// Response is nil by default
	mp = ""

	// Load the pattern length
	n := len(pattern)

	// No empty patterns
	if n == 0 {
		return
	}

	// Patterns should have a trailing slash in order to match anyting
	if pattern[n-1] != '/' {

		// Patter is different
		if pattern == path {
			mp = "/"
		}

		// Return
		return
	}

	// Check if the pattern matches
	if len(path) >= n && path[0:n] == pattern {
		mp = path[n-1:]
	}

	return
}
