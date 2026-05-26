// Package gomponents provides HTML components in Go, that render to HTML 5.
//
// The primary interface is a [Node]. It defines a function Render, which should render the [Node]
// to the given writer as a string.
//
// All DOM elements and attributes can be created by using the [El] and [Attr] functions.
//
// The functions [Text], [Textf], [Raw], and [Rawf] can be used to create text nodes, either HTML-escaped or unescaped.
//
// See also helper functions [Map], [If], and [Iff] for mapping data to nodes and inserting them conditionally.
//
// There's also the [Group] type, which is a slice of [Node]-s that can be rendered as one [Node].
//
// For basic HTML elements and attributes, see the package html.
//
// For higher-level HTML components, see the package components.
//
// For HTTP helpers, see the package http.
package gomponents

import (
	"fmt"
	"io"
)

// Node is a DOM node that can Render itself to a [io.Writer].
type Node interface {
	Render(w io.Writer) error
}

// NodeType describes what type of [Node] it is, currently either an [ElementType] or an [AttributeType].
// This decides where a [Node] should be rendered.
// Nodes default to being [ElementType].
type NodeType int

const (
	ElementType = NodeType(iota)
	AttributeType
)

// nodeTypeDescriber can be implemented by Nodes to let callers know whether the [Node] is
// an [ElementType] or an [AttributeType].
// See [NodeType].
type nodeTypeDescriber interface {
	Type() NodeType
}

// Compile-time check that [NodeFunc] implements [fmt.Stringer], [Node] and [nodeTypeDescriber].
var _ interface {
	fmt.Stringer
	Node
	nodeTypeDescriber
} = (NodeFunc)(nil)

// NodeFunc is a render function that is also a [Node] of [ElementType].
type NodeFunc func(io.Writer) error

// Render satisfies [Node].
func (n NodeFunc) Render(w io.Writer) error {
	_ = "STUB: not implemented"

	// Type satisfies [nodeTypeDescriber].
	return nil
}

func (NodeFunc) Type() NodeType {
	_ = "STUB: not implemented"

	// String satisfies [fmt.Stringer].
	return *new(NodeType)
}

func (n NodeFunc) String() string { _ = "STUB: not implemented"; return "" }

var (
	lt      = []byte("<")
	gt      = []byte(">")
	ltSlash = []byte("</")
)

// El creates an element DOM [Node] with a name and child Nodes.
// See https://dev.w3.org/html5/spec-LC/syntax.html#elements-0 for how elements are rendered.
// No tags are ever omitted from normal tags, even though it's allowed for elements given at
// https://dev.w3.org/html5/spec-LC/syntax.html#optional-tags
// If an element is a void element, non-attribute children nodes are ignored.
// Use this if no convenience creator exists in the html package.
func El(name string, children ...Node) Node { _ = "STUB: not implemented"; return *new(Node) }

// renderChild c to the given writer w if the node type is desiredType.
func renderChild(w io.Writer, c Node, desiredType NodeType) error {
	_ = "STUB: not implemented"
	return nil
}

// Rendering groups like this is still important even though a group can render itself,
// since otherwise attributes will sometimes be ignored.

// isVoidElement reports whether the named element is a void element that doesn't have an end tag.
// See https://dev.w3.org/html5/spec-LC/syntax.html#void-elements
func isVoidElement(name string) bool { _ = "STUB: not implemented"; return false }

var (
	space      = []byte(" ")
	equalQuote = []byte(`="`)
	quote      = []byte(`"`)
)

// Attr creates an attribute DOM [Node] with a name and optional value.
// If only a name is passed, it's a name-only (boolean) attribute (like "required").
// If a name and value are passed, it's a name-value attribute (like `class="header"`).
// More than one value makes [Attr] panic.
// Use this if no convenience creator exists in the html package.
func Attr(name string, value ...string) Node { _ = "STUB: not implemented"; return *new(Node) }

// booleanAttr creates a boolean attribute Node with just a name.
func booleanAttr(name string) Node { _ = "STUB: not implemented"; return *new(Node) }

// valueAttr creates a name-value attribute Node.
func valueAttr(name, value string) Node { _ = "STUB: not implemented"; return *new(Node) }

// Compile-time check that [attrFunc] implements [fmt.Stringer], [Node] and [nodeTypeDescriber].
var _ interface {
	fmt.Stringer
	Node
	nodeTypeDescriber
} = (attrFunc)(nil)

// attrFunc is a render function that is also a [Node] of [AttributeType].
// It's basically the same as [NodeFunc], but for attributes.
type attrFunc func(io.Writer) error

// Render satisfies [Node].
func (a attrFunc) Render(w io.Writer) error {
	_ = "STUB: not implemented"

	// Type satisfies [nodeTypeDescriber].
	return nil
}

func (attrFunc) Type() NodeType {
	_ = "STUB: not implemented"
	return *

	// String satisfies [fmt.Stringer].
	new(NodeType)
}

func (a attrFunc) String() string { _ = "STUB: not implemented"; return "" }

// Text creates a text DOM [Node] that Renders the escaped string t.
func Text(t string) Node { _ = "STUB: not implemented"; return *new(Node) }

// Textf creates a text DOM [Node] that Renders the interpolated and escaped string format.
func Textf(format string, a ...interface{}) Node { _ = "STUB: not implemented"; return *new(Node) }

// Compile-time check that [raw] implements [fmt.Stringer], [Node], and [nodeTypeDescriber].
var _ interface {
	fmt.Stringer
	Node
	nodeTypeDescriber
} = raw("")

// raw is a text DOM [Node] that just Renders the unescaped, underlying string.
type raw string

func (r raw) Render(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (r raw) String() string { _ = "STUB: not implemented"; return "" }

func (r raw) Type() NodeType {
	_ = "STUB: not implemented"

	// Raw creates a text DOM [Node] that just Renders the unescaped string t.
	return *new(NodeType)
}

func Raw(t string) Node {
	_ = "STUB: not implemented"

	// Rawf creates a text DOM [Node] that just Renders the interpolated and unescaped string format.
	return *new(Node)
}

func Rawf(format string, a ...interface{}) Node { _ = "STUB: not implemented"; return *new(Node) }

// Map a slice of anything to a [Group] (which is just a slice of [Node]-s).
func Map[T any](ts []T, cb func(T) Node) Group { _ = "STUB: not implemented"; return *new(Group) }

// Compile-time check that [Group] implements [fmt.Stringer] and [Node].
var _ interface {
	fmt.Stringer
	Node
} = (Group)(nil)

// Group a slice of [Node]-s into one Node, while still being usable like a regular slice of [Node]-s.
// A [Group] can render directly, but if any of the direct children are [AttributeType], they will be ignored,
// to not produce invalid HTML.
type Group []Node

// String satisfies [fmt.Stringer].
func (g Group) String() string { _ = "STUB: not implemented"; return "" }

// Render satisfies [Node].
func (g Group) Render(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// If condition is true, return the given [Node]. Otherwise, return nil.
// This helper function is good for inlining elements conditionally.
// If it's important that the given [Node] is only evaluated if condition is true
// (for example, when using nilable variables), use [Iff] instead.
func If(condition bool, n Node) Node { _ = "STUB: not implemented"; return *new(Node) }

// Iff condition is true, call the given function. Otherwise, return nil.
// This helper function is good for inlining elements conditionally when the node depends on nilable data,
// or some other code that could potentially panic.
// If you just need simple conditional rendering, see [If].
func Iff(condition bool, f func() Node) Node { _ = "STUB: not implemented"; return *new(Node) }
