package main

// actionOutputMap is a map of string -> action were the key is the primary
// output file(s) of the action.
type actionOutputMap map[string]*Action

// newActionOutputMap creates a new actionOutputMap.
func newActionOutputMap(actions []*Action) actionOutputMap {
	result := make(actionOutputMap)
	for _, action := range actions {
		result[action.OutputFiles] = action
	}
	return result
}
