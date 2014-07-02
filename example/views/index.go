package views

import (
	"github.com/mgutz/gorazor/example/models"
	"github.com/mgutz/gorazor/example/views/layout"
	"github.com/mgutz/gorazor/gorazor"
)

func Index(user *models.User) gorazor.SafeBuffer {
	_buffer := gorazor.NewSafeBuffer()
	_buffer.WriteString("\n\n<p>Escaped: ")
	_buffer.WriteSafe(UnsafeHello(user.Name))
	_buffer.WriteString("</p>\n<p>Unescaped: ")
	_buffer.WriteSafe(SafeHello(user.Name))
	_buffer.WriteString("</p>")
	_buffer.WriteSafe(Raw("<h2>Heading 2</h2>"))

	js := func() gorazor.SafeBuffer {
		_buffer := gorazor.NewSafeBuffer()

		_buffer.WriteString("<script>\n  alert('Hello, ")
		_buffer.WriteSafe(user.Name)
		_buffer.WriteString("');\n</script>")

		return _buffer
	}

	title := func() gorazor.SafeBuffer {
		_buffer := gorazor.NewSafeBuffer()

		_buffer.WriteString("\"Cool Site\"")
		return _buffer
	}

	return layout.Default(_buffer, js(), title())
}
