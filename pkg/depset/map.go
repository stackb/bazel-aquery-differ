package depset

import anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"

// Map is a map of id -> depSetOfFiles.
type Map map[uint32]*anpb.DepSetOfFiles

// NewMap creates a new map of depSet by its ID.
func NewMap(depSets []*anpb.DepSetOfFiles) Map {
	result := make(Map)
	for _, depSet := range depSets {
		result[depSet.Id] = depSet
	}
	return result
}
