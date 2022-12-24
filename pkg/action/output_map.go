package action

import (
	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
)

// OutputMap is a map of string -> action were the key is the primary
// output file(s) of the action.
type OutputMap map[string]*dipb.Action

// NewOutputMap creates a new actionOutputMap.
func NewOutputMap(actions []*dipb.Action) OutputMap {
	result := make(OutputMap)
	for _, action := range actions {
		result[action.OutputFiles] = action
	}
	return result
}
