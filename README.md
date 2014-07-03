# razor

`razor` is the Go port of the ASP.NET's Razor view engine.

`razor` is a code generation library which converts Razor templates into Go functions.

# Usage

Install

```sh
go get github.com/mgutz/razor
```

Running

```sh
razor template_folder output_folder
razor template_file output_file
```

## Layout & Views

Let's cover the basic case of a view with a layout. In `razor` templates become
functions.

A layout is nothing more than function which receives the rendered result of a view.
That is, given a layout function named `Layout` and a view function `View`, the view
is rendered as `Layout(View())`.

Let's see this in action.  First define a layout, `views/layout/base.gohtml`

```html
@meta {
    +func(title string, css razor.SafeBuffer, body razor.SafeBuffer, js razor.SafeBuffer)
}

<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8" />
	<title>@title</title>
        @css
</head>
<body>
        <div class="container">@body</div>
        @js
</body>
</html>
```

The first block of template instructs  `razor` how to generate the function.  The
layout declares a function with a signature of

    (title string, css razor.SafeBuffer, body razor.SafeBuffer, js razor.SafeBuffer)

The arguments are used in the template body and denoted with `@`.

Let's now define a view `views/index.gohtml` to use the layout.

```html
@meta {
    import (
        "views/layout"
    )

    +func(name string)
    +return layout.Base(title, "", VIEW, js())
}

@{
    // inline code. You know about code separation of concerns right?
    title := "Welcome Page"
}

<h2>Welcome to homepage</h2>

@section js {
<script>
    alert('hello! @name')
</script>
}
```

This view has a signature of `(name string)` which means `name` will be passed in
as an argument.  A variable `title` is set in a code block and is used by the layout.
A section named `js` becomes its own function. The magic all happens in the
function's return value of `layout.Base(title, "", VIEW, js())`. `VIEW` is a placeholder
for the rendered value of the view template.

To call from Go code

```go
import (
    "views"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
    views.Index("gopherito").WriteTo(w)
}
```

See [working example](example).

| Description | Template | Generated code |
| ------------| -------- | ---------------|
| View |  [index.gohtml](example/views/index.gohtml) | [index.go](example/views/index.go) |
| Layout | [default.gohtml](example/views/layout/default.gohtml) | [default.go](example/views/layout/default.go) |

## Syntax

### Variable

* `@variable` to insert **string** variable into html template
* variable could be wrapped by arbitrary go functions
* variable inserted will be automatically [escaped](http://golang.org/pkg/html/template/#HTMLEscapeString)

```html
<div>Hello @user.Name</div>
```

```html
<div>Hello @strings.ToUpper(req.CurrentUser.Name)</div>
```

`razor` escapes any value that is not of type `razor.SafeBuffer`. To
insert unescaped data create a helper function. See [example helper](views/helper.go) directory:

```go
func Raw(t interface{}) gorazor.SafeBuffer {
	// Safe = true tells `gorazor` this buffer is safe to write as-is
	buffer := gorazor.NewSafeBuffer()
	buffer.WriteString(fmt.Sprint(t))
	return buffer
}
```


### Flow Control

```php
@if .... {
	....
}

@if .... {
	....
} else {
	....
}

@for .... {

}

@{
	switch .... {
	case ....:
	      <p>...</p>
	case 2:
	      <p>...</p>
	default:
	      <p>...</p>
	}
}
```

### Code block

It's possible to insert arbitrary go code block in the template, like create new variable.

```html
@{
	username := u.Name
	if u.Email != "" {
		username += "(" + u.Email + ")"
	}
}
<div class="welcome">
<h4>Hello @username</h4>
</div>
```

### Declaration

The **first code block** in the template is strictly for declaration:

* imports
* function signature
* layout return value

For example:

```go
@meta {
import  (
    "kp/models"
    "kp/views/layout"
)

// Generated function signature
+func(user *models.User, blog *models.Blog)

// Override the return value to call another function (used for layouts).
+return layout.Default(VIEW, section1(), section2())
}
...
```

Results in a view method:

```go
package dirname

import (
    "kp/models"
    "kp/views/layout"
)

func Basename(user *models.User, blog *models.Blog) razor.SafeBuffer {
    _buffer := razor.NewSafeBuffer()
    _buffer := layout.Default(_buffer, section1(), section2())
    return _buffer
}
```

**first code block** must be at the beginning of the template, i.e. before any html.

Any other codes inside the first code block will **be ignored**.

`import` must be wrapped in `()`

The variables declared in **first code block** will be the models of the template, i.e. the parameters of generated function.

If your template doesn't need any model input, then just leave it blank.


### Helper / Include other template

`razor` compiles templates to go function, embedding another template is
just calling another go function.


* A layout should be able to use another layout, it's just function call.

## Conventions

* Template **folder name** will be used as **package name** in generated code
* Template file name must have the extension name `.gohtml`
* The **function name** is the Capitalized basename of the file (without extension).


## FAQ

## How to auto re-generate when gohtml file changes?

Use the right tool for the job. I recommend [node.js](https://nodejs.org) and
[gulp](https://gulpjs.com) at this time. As of now build and asset preprocessing is
lacking for gophers. Gulp also supports such things as LESS, SASS, minification,
optimizing images, Common JS and is fast ...

See `example` directory for an example `gulpfile`

# Credits

The original work [sipin's gorazor](https://github.com/sipin/gorazor)

