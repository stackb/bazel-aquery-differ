package action

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"google.golang.org/protobuf/proto"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
	"github.com/stackb/bazel-aquery-differ/pkg/protobuf"
)

type OutputPair struct {
	Output string
	Action *dipb.Action // representative of before/after
	Before *dipb.Action
	After  *dipb.Action
}

func (p *OutputPair) Diff() string {
	return cmp.Diff(
		p.Before,
		p.After,
		cmpopts.IgnoreFields(dipb.Action{}, "Id"),
		cmpopts.IgnoreUnexported(dipb.Action{}, anpb.KeyValuePair{}),
	)
}

func (p *OutputPair) UnifiedDiff() gotextdiff.Unified {
	var a string
	var b string
	if p.Before != nil {
		beforeCopy := proto.Clone(p.Before).(*dipb.Action)
		beforeCopy.Id = ""
		a = protobuf.FormatProtoText(beforeCopy)
	}
	if p.After != nil {
		afterCopy := proto.Clone(p.After).(*dipb.Action)
		afterCopy.Id = ""
		b = protobuf.FormatProtoText(afterCopy)
	}
	edits := myers.ComputeEdits(span.URI(p.Output), a, b)
	return gotextdiff.ToUnified(p.Output, p.Output, a, edits)
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
