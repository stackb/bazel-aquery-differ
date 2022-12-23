package main

import (
	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type actionOutputPair struct {
	output string
	before *Action
	after  *Action
}

func (p *actionOutputPair) diff() string {
	return cmp.Diff(p.before, p.after, cmpopts.IgnoreUnexported(anpb.KeyValuePair{}))
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
