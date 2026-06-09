package routers

import (
	"testing"
)

// TestRouterInitDoesNotPanic verifies that the init() function
// registers routes without panicking
func TestRouterInitDoesNotPanic(t *testing.T) {
	// init() is called automatically when the package is imported.
	// If we reach this line, router registration succeeded.
	t.Log("Router init completed without panic")
}
