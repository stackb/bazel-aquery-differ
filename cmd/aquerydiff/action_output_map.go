package main

import (
	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
)

// actionOutputMap is a map of string -> action were the key is the primary
// output file(s) of the action.
type actionOutputMap map[string]*dipb.Action

// newActionOutputMap creates a new actionOutputMap.
func newActionOutputMap(actions []*dipb.Action) actionOutputMap {
	result := make(actionOutputMap)
	for _, action := range actions {
		result[action.OutputFiles] = action
	}
	return result
}
