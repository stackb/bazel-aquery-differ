package report

import "github.com/stackb/bazel-aquery-differ/pkg/action"

type Html struct {
	Dir        string
	BeforeOnly action.OutputPairs
	AfterOnly  action.OutputPairs
	Equal      action.OutputPairs
	Different  action.OutputPairs
}
