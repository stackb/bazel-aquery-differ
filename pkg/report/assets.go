package report

import (
	"embed"
)

//go:embed index.html.tmpl
var indexHtmlFs embed.FS

//go:embed diff.html.tmpl
var diffHtmlFs embed.FS

//go:embed cmp.html.tmpl
var cmpHtmlFs embed.FS

//go:embed style.css
var styleCss []byte
