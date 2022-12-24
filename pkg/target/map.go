package target

import (
	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
)

// targetMap is a map of id -> target.
type Map map[uint32]*anpb.Target

// NewMap creates new map of targets by ID.
func NewMap(targets []*anpb.Target) Map {
	result := make(Map)
	for _, target := range targets {
		result[target.Id] = target
	}
	return result
}
