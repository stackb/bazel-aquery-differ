package report

import (
	"embed"
)

//go:embed index.html.tmpl
var indexHtmlFs embed.FS

//go:embed style.css
var styleCss []byte
