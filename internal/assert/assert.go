// Package assert provides testing helpers.
package assert

import (
	"testing"

	g "maragu.dev/gomponents"
)

// Equal checks for equality between the given expected string and the rendered Node string.
func Equal(t *testing.T, expected string, actual g.Node) { _ = "STUB: not implemented"; return }

// Error checks for a non-nil error.
func Error(t *testing.T, err error) { _ = "STUB: not implemented"; return }
