package gorazor

import (
	"bytes"
	"html/template"
	"io"
)

type SafeBuffer struct {
	bytes.Buffer
	Safe bool
}

func (self *SafeBuffer) WriteTo(w io.Writer) {
	if self.Safe {
		self.Buffer.WriteTo(w)
	} else {
		template.HTMLEscape(w, self.Buffer.Bytes())
	}
}

func (self *SafeBuffer) WriteSafe(t interface{}) {
	switch v := t.(type) {
	case *SafeBuffer:
		if v.Safe {
			self.Write(v.Bytes())
		} else {
			self.WriteString(template.HTMLEscapeString(v.String()))
		}
	default:
		self.WriteString(template.HTMLEscaper(v))
	}
}
