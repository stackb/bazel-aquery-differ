package main

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pmezard/go-difflib/difflib"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
)

type actionOutputPair struct {
	output string
	before *dipb.Action
	after  *dipb.Action
}

func (p *actionOutputPair) diff() string {
	return cmp.Diff(p.before, p.after, cmpopts.IgnoreUnexported(anpb.KeyValuePair{}))
}

func (p *actionOutputPair) unifiedDiff(fromFile, toFile string) string {
	var a string
	var b string
	if p.before != nil {
		a = FormatAction(p.before)
	}
	if p.after != nil {
		b = FormatAction(p.after)
	}
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(a),
		B:        difflib.SplitLines(b),
		FromFile: fromFile,
		ToFile:   toFile,
		Context:  3,
	}
	text, _ := difflib.GetUnifiedDiffString(diff)
	return text
}

type actionOutputPairs []*actionOutputPair

func (p actionOutputPairs) Len() int {
	return len(p)
}

func (p actionOutputPairs) Less(i, j int) bool {
	a := p[i]
	b := p[j]
	return a.output < b.output
}

func (p actionOutputPairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
