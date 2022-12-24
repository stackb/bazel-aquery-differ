package report

import (
	"io"
	"os"
	"path/filepath"

	"html/template"

	"github.com/stackb/bazel-aquery-differ/pkg/action"
)

type Html struct {
	BeforeFile string
	AfterFile  string
	BeforeOnly action.OutputPairs
	AfterOnly  action.OutputPairs
	Equal      action.OutputPairs
	Different  action.OutputPairs
}

func (r *Html) Emit(dir string) error {
	if err := r.emitIndex(dir); err != nil {
		return err
	}
	return nil
}

func (r *Html) emitIndex(dir string) error {
	filename := filepath.Join(dir, "index.html")
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	return renderIndexHtml(out, r.BeforeFile, r.AfterFile)
}

func renderIndexHtml(out io.Writer, beforeFile, afterFile string) error {
	tmpl := template.Must(template.New("index.html.tmpl").ParseFS(indexHtmlFs, "index.html.tmpl"))
	data := struct {
		BeforeFile, AfterFile string
	}{
		BeforeFile: beforeFile,
		AfterFile:  afterFile,
	}
	return tmpl.Execute(out, data)
}
