// DO NOT EDIT! Auto-generated by github.com/mgutz/gorazor
package cases

import (
	"github.com/mgutz/gorazor/razor"
)

func Quote() razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	_buffer.WriteString("<html>'text'</html>")

	return _buffer
}
