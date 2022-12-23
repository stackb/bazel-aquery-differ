package main

import (
	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
)

// targetMap is a map of id -> target.
type targetMap map[uint32]*anpb.Target

// newTargetMap creates new map of targets by ID.
func newTargetMap(targets []*anpb.Target) targetMap {
	result := make(targetMap)
	for _, target := range targets {
		result[target.Id] = target
	}
	return result
}
