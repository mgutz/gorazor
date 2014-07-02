package layout

import (
	"github.com/mgutz/gorazor/gorazor"
)

func Default(body *gorazor.SafeBuffer, js *gorazor.SafeBuffer, title *gorazor.SafeBuffer) *gorazor.SafeBuffer {
	_buffer := &gorazor.SafeBuffer{Safe: true}
	_buffer.WriteString("\n\n<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\" />\n    <title>")
	_buffer.WriteSafe(title)
	_buffer.WriteString("</title>\n</head>\n<body>\n    <div class=\"container\">")
	_buffer.WriteSafe(body)
	_buffer.WriteString("</div>\n    ")
	_buffer.WriteSafe(js)
	_buffer.WriteString("\n  </body>\n</html>")

	return _buffer
}
