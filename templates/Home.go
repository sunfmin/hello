package templates

import (
	"bytes"
	"fmt"
	"github.com/sipin/gorazor/gorazor"
)

func Home() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<p> ")
	_buffer.WriteString(gorazor.HTMLEscape(fmt.Sprintf("This Go generated part: %s", "Yeah~~")))
	_buffer.WriteString(" </p>\n\nThis is react generated part\n<div id=\"root\"></div>")

	return _buffer.String()
}
