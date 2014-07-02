package views

import (
	"fmt"
	"github.com/mgutz/gorazor/gorazor"
)

// This will be escaped by the template
func UnsafeHello(name string) string {
	return "Hello <i>" + name + "</i>!"
}

// Will not be escaped.
func SafeHello(name string) *gorazor.SafeBuffer {
	// Safe = true tells `gorazor` this buffer is safe to write as-is
	buffer := &gorazor.SafeBuffer{Safe: true}

	buffer.WriteString("Hello <i>")
	buffer.WriteSafe(name)
	buffer.WriteString("</i>!")
	return buffer
}

func Raw(t interface{}) *gorazor.SafeBuffer {
	// Safe = true tells `gorazor` this buffer is safe to write as-is
	buffer := &gorazor.SafeBuffer{Safe: true}
	buffer.WriteString(fmt.Sprint(t))
	return buffer
}
