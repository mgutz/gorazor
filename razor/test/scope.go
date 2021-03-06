// DO NOT EDIT! Auto-generated by github.com/mgutz/gorazor
package cases

import (
	"dm"
	"github.com/mgutz/gorazor/razor"
	"zfw/models"
	. "zfw/tplhelper"
)

func Scope() razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	var obj *models.Widget
	{
		data, dmType := dm.GetData(obj.PlaceHolder)

		if dmType == "simple" {
			obj.StringList = data.([]string)

			_buffer.WriteString("<div>")
			_buffer.WriteSafe((SelectPk(obj)))
			_buffer.WriteString("</div>")

		} else {
			node := data.(*dm.DMTree)

			_buffer.WriteString("<div class=\"form-group ")
			_buffer.WriteSafe(GetErrorClass(obj))
			_buffer.WriteString("\">\n  <label for=\"")
			_buffer.WriteSafe(obj.Name)
			_buffer.WriteString("\" class=\"col-sm-2 control-label\">")
			_buffer.WriteSafe(obj.Label)
			_buffer.WriteString("</label>\n  <div class=\"col-sm-10\">\n    <select class=\"form-control\" name=\"")
			_buffer.WriteSafe(obj.Name)
			_buffer.WriteString("\" ")
			_buffer.WriteSafe(BoolStr(obj.Disabled, "disabled"))
			_buffer.WriteString(">\n      ")
			for _, option := range node.Keys {
				if values, ok := node.Values[option]; ok {

					_buffer.WriteString("<optgroup label=\"")
					_buffer.WriteSafe(option)
					_buffer.WriteString("\">\n        ")
					for _, value := range values {
						if value == obj.Value {

							_buffer.WriteString("<option selected>")
							_buffer.WriteSafe(value)
							_buffer.WriteString("</option>")

						} else {

							_buffer.WriteString("<option>")
							_buffer.WriteSafe(value)
							_buffer.WriteString("</option>")

						}
					}
					_buffer.WriteString("\n      </optgroup>")

				} else {
					if option == obj.Value {

						_buffer.WriteString("<option selected>")
						_buffer.WriteSafe(option)
						_buffer.WriteString("</option>")

					} else {

						_buffer.WriteString("<option>")
						_buffer.WriteSafe(option)
						_buffer.WriteString("</option>")

					}
				}
			}
			_buffer.WriteString("\n    </select>\n    ")
			if obj.ErrorMsg != "" {

				_buffer.WriteString("<span class=\"label label-danger\">")
				_buffer.WriteSafe(obj.ErrorMsg)
				_buffer.WriteString("</span>")

			}
			_buffer.WriteString("\n  </div>\n</div>")
		}
	}

	return _buffer
}
