package report

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"html/template"

	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
	"github.com/stackb/bazel-aquery-differ/pkg/action"
	"github.com/stackb/bazel-aquery-differ/pkg/protobuf"
)

type Html struct {
	BeforeFile string
	AfterFile  string
	Before     *action.Graph
	After      *action.Graph
	BeforeOnly action.OutputPairs
	AfterOnly  action.OutputPairs
	Equal      action.OutputPairs
	NonEqual   action.OutputPairs
}

func (r *Html) Emit(dir string) error {
	if err := r.emitFile(dir, r.BeforeFile); err != nil {
		return err
	}
	r.BeforeFile = filepath.Base(r.BeforeFile)

	if err := r.emitFile(dir, r.AfterFile); err != nil {
		return err
	}
	r.AfterFile = filepath.Base(r.AfterFile)

	for _, action := range r.Before.Actions {
		if err := r.emitAction(dir, action); err != nil {
			return err
		}
	}

	for _, action := range r.After.Actions {
		if err := r.emitAction(dir, action); err != nil {
			return err
		}
	}

	for _, pair := range r.Equal {
		if err := r.emitOutputPairDiff(dir, pair); err != nil {
			return err
		}
	}

	for _, pair := range r.NonEqual {
		if err := r.emitOutputPairDiff(dir, pair); err != nil {
			return err
		}
	}

	if err := r.emitIndexHtml(dir); err != nil {
		return err
	}

	if err := r.emitStyleCss(dir); err != nil {
		return err
	}

	return nil
}

func (r *Html) emitAction(dir string, action *dipb.Action) error {
	if err := r.emitActionJsonproto(dir, action); err != nil {
		return err
	}
	if err := r.emitActionTextproto(dir, action); err != nil {
		return err
	}
	return nil
}

func (r *Html) emitFile(dir string, original string) error {
	filename := filepath.Join(dir, filepath.Base(original))
	basedir := filepath.Dir(filename)
	if err := os.MkdirAll(basedir, os.ModePerm); err != nil {
		return err
	}
	return copyFile(original, filename)
}

func (r *Html) emitActionJsonproto(dir string, action *dipb.Action) error {
	filename := filepath.Join(dir, action.Id+".json")
	basedir := filepath.Dir(filename)
	if err := os.MkdirAll(basedir, os.ModePerm); err != nil {
		return err
	}
	if err := protobuf.WritePrettyJSONFile(filename, action); err != nil {
		return err
	}
	return nil
}

func (r *Html) emitActionTextproto(dir string, a *dipb.Action) error {
	filename := filepath.Join(dir, a.Id+".textproto")
	basedir := filepath.Dir(filename)
	if err := os.MkdirAll(basedir, os.ModePerm); err != nil {
		return err
	}
	if err := protobuf.WritePrettyTextFile(filename, a); err != nil {
		return err
	}
	return nil
}

func (r *Html) emitOutputPairDiff(dir string, pair *action.OutputPair) error {
	filename := filepath.Join(dir, pair.Before.Id, pair.After.Id)
	basedir := filepath.Dir(filename)
	if err := os.MkdirAll(basedir, os.ModePerm); err != nil {
		return err
	}
	if err := os.WriteFile(filename+".diff.txt", []byte(pair.UnifiedDiff()), fs.ModePerm); err != nil {
		return err
	}
	if err := os.WriteFile(filename+".cmp.txt", []byte(pair.Diff()), fs.ModePerm); err != nil {
		return err
	}
	return nil
}

func (r *Html) emitIndexHtml(dir string) error {
	filename := filepath.Join(dir, "index.html")
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	return r.renderIndexHtml(out)
}

func (r *Html) emitStyleCss(dir string) error {
	filename := filepath.Join(dir, "style.css")
	return ioutil.WriteFile(filename, styleCss, os.ModePerm)
}

func (r *Html) renderIndexHtml(out io.Writer) error {
	tmpl := template.Must(template.New("index.html.tmpl").ParseFS(indexHtmlFs, "index.html.tmpl"))
	return tmpl.Execute(out, r)
}
