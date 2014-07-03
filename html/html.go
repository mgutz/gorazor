package html

import (
	"bytes"
	"fmt"

	"github.com/mgutz/gorazor/razor"
)

func Raw(t interface{}) razor.SafeBuffer {
	buffer := razor.NewSafeBuffer()
	switch v := t.(type) {
	case razor.SafeBuffer:
		buffer.Write(v.Bytes())
	case bytes.Buffer:
		buffer.Write(v.Bytes())
	default:
		buffer.WriteString(fmt.Sprint(t))
	}
	return buffer
}
