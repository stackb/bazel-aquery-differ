package action

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pmezard/go-difflib/difflib"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
)

type OutputPair struct {
	Output string
	Before *dipb.Action
	After  *dipb.Action
}

func (p *OutputPair) Diff() string {
	return cmp.Diff(p.Before, p.After, cmpopts.IgnoreUnexported(dipb.Action{}, anpb.KeyValuePair{}))
}

func (p *OutputPair) UnifiedDiff() string {
	var a string
	var b string
	if p.Before != nil {
		a = FormatAction(p.Before)
	}
	if p.After != nil {
		b = FormatAction(p.After)
	}
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(a),
		B:        difflib.SplitLines(b),
		FromFile: p.Output,
		ToFile:   p.Output,
		Context:  3,
	}
	text, _ := difflib.GetUnifiedDiffString(diff)
	return text
}

type OutputPairs []*OutputPair

func (p OutputPairs) Len() int {
	return len(p)
}

func (p OutputPairs) Less(i, j int) bool {
	a := p[i]
	b := p[j]
	return a.Output < b.Output
}

func (p OutputPairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
