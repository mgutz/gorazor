package views

import (
	"github.com/mgutz/gorazor/razor"
)

// This will be escaped by the template
func UnsafeHello(name string) string {
	return "Hello <i>" + name + "</i>!"
}

// Will not be escaped since we are using SafeBuffer.
func SafeHello(name string) razor.SafeBuffer {
	buffer := razor.NewSafeBuffer()
	buffer.WriteString("Hello <i>")
	buffer.WriteSafe(name)
	buffer.WriteString("</i>!")
	return buffer
}
