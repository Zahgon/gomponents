// Package components provides high-level components and helpers that are composed of low-level elements and attributes.
package components

import (
	"io"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// HTML5Props for [HTML5].
// Title is set no matter what, Description and Language elements only if the strings are non-empty.
type HTML5Props struct {
	Title       string
	Description string
	Language    string
	Head        g.Group
	Body        g.Group
	HTMLAttrs   g.Group
}

// HTML5 document template.
func HTML5(p HTML5Props) g.Node { _ = "STUB: not implemented"; return *new(g.Node) }

// Classes is a map of strings to booleans, which Renders to an attribute with name "class".
// The attribute value is a sorted, space-separated string of all the map keys,
// for which the corresponding map value is true.
type Classes map[string]bool

// Render satisfies [g.Node].
func (c Classes) Render(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (c Classes) Type() g.NodeType {
	_ = "STUB: not implemented"
	return *

	// String satisfies [fmt.Stringer].
	new(g.NodeType)
}

func (c Classes) String() string { _ = "STUB: not implemented"; return "" }

// JoinAttrs joins attributes with the given name on the first level of the given nodes.
// Attributes on non-direct descendants are ignored.
// Non-empty attribute values are joined by spaces into a single attribute.
// Empty, whitespace-only, and boolean (valueless) attributes are deduplicated and discarded
// if a non-empty value exists. If only boolean/empty attributes match, a single boolean
// attribute is emitted.
// When both boolean and valued attributes match, the valued form takes precedence.
// Note that this renders all first-level attributes to check whether they should be processed.
func JoinAttrs(name string, children ...g.Node) g.Node {
	_ = "STUB: not implemented"
	return *new(g.Node)
}

// processNode checks a single child node and either collects its value or appends it to result.

// If no matching attributes were found, just return the result now

// Insert joined attribute at the position of the first match

type nodeTypeDescriber interface {
	Type() g.NodeType
}

func extractAttrValue(name string, n g.Node) (bool, string) {
	_ = "STUB: not implemented"
	// Ignore everything that is not an attribute
	return false, ""
}

// Match boolean attribute (e.g., ` required`)

// Unescape to get the original value, since it will be escaped again when the joined attribute is rendered

// Treat whitespace-only values the same as empty
