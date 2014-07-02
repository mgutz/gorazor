package gorazor

import (
	"bytes"
	"html/template"
	"io"
)

type SafeBuffer struct {
	*bytes.Buffer
}

func NewSafeBuffer() SafeBuffer {
	return SafeBuffer{Buffer: bytes.NewBuffer(nil)}
}

func (self SafeBuffer) WriteTo(w io.Writer) {
	self.Buffer.WriteTo(w)
}

func (self SafeBuffer) WriteSafe(t interface{}) {
	switch v := t.(type) {
	case SafeBuffer:
		self.Write(v.Bytes())
	default:
		self.WriteString(template.HTMLEscaper(v))
	}
}
