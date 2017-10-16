package templates

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	m "github.com/theplant/createreactappmanifest"
)

func Layout(mani *m.Manifest, content string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<!DOCTYPE html>\n<html>\n  <head>\n    <meta charset=\"utf-8\">\n    <title>My page</title>\n    <link href=\"")
	_buffer.WriteString(gorazor.HTMLEscape(mani.GetURL("main.css")))
	_buffer.WriteString("\" rel=\"stylesheet\">\n  </head>\n  <body>\n    ")
	_buffer.WriteString((content))
	_buffer.WriteString("\n    <script type=\"text/javascript\" src=\"")
	_buffer.WriteString(gorazor.HTMLEscape(mani.GetURL("main.js")))
	_buffer.WriteString("\"></script>\n  </body>\n</html>")

	return _buffer.String()
}
