package main

import anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"

// depSetOfFilesMap is a map of id -> depSetOfFiles.
type depSetOfFilesMap map[uint32]*anpb.DepSetOfFiles

// newDepSetOfFilesMap creates a new map of depSet by its ID.
func newDepSetOfFilesMap(depSets []*anpb.DepSetOfFiles) depSetOfFilesMap {
	result := make(depSetOfFilesMap)
	for _, depSet := range depSets {
		result[depSet.Id] = depSet
	}
	return result
}
