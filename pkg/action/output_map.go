package action

import (
	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
)

// ActionMap is a map of string -> action.
type ActionMap map[string]*dipb.Action

// ActionMapper is a function that creates an ActionMap from a list of Actions.
type ActionMapper func([]*dipb.Action) ActionMap

// NewOutputFilesMap creates a new actionOutputMap were the key is the primary
// output file(s) of the action.
func NewOutputFilesMap(actions []*dipb.Action) ActionMap {
	result := make(ActionMap)
	for _, action := range actions {
		result[action.OutputFiles] = action
	}
	return result
}

// NewMnemonicFileMap creates a new actionOutputMap were the key is the primary
// output file(s) of the action.  In the strategy, the last action having the mnemonic wins.
func NewMnemonicFileMap(actions []*dipb.Action) ActionMap {
	result := make(ActionMap)
	for _, action := range actions {
		result[action.Mnemonic] = action
	}
	return result
}
