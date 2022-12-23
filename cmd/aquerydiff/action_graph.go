package main

import (
	"sort"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
)

// actionGraph holds compiled data about the action graph container
type actionGraph struct {
	container       *anpb.ActionGraphContainer
	artifacts       artifactPathMap
	targets         targetMap
	depSetOfFiles   depSetOfFilesMap
	depSetResolver  depSetResolver
	actions         []*Action
	actionOutputMap actionOutputMap
}

func newActionGraph(container *anpb.ActionGraphContainer) (*actionGraph, error) {
	artifacts, err := newArtifactPathMap(container.Artifacts, container.PathFragments)
	if err != nil {
		return nil, err
	}

	targets := newTargetMap(container.Targets)
	depSetOfFiles := newDepSetOfFilesMap(container.DepSetOfFiles)
	depSetResolver := newDepSetResolver(artifacts, depSetOfFiles)

	actions := make([]*Action, len(container.Actions))
	for i, a := range container.Actions {
		action, err := newAction(a, artifacts, targets, *depSetResolver)
		if err != nil {
			return nil, err
		}
		actions[i] = action
	}

	return &actionGraph{
		container:       container,
		artifacts:       artifacts,
		targets:         targets,
		depSetOfFiles:   depSetOfFiles,
		depSetResolver:  *depSetResolver,
		actions:         actions,
		actionOutputMap: newActionOutputMap(actions),
	}, nil
}

func partitionGraphs(before, after actionOutputMap) (beforeOnly, afterOnly, both actionOutputPairs) {
	a := make(map[string]bool)
	b := make(map[string]bool)
	for output := range before {
		a[output] = true
	}
	for output := range after {
		b[output] = true
	}
	for output := range a {
		if b[output] {
			both = append(both, &actionOutputPair{output: output, before: before[output], after: after[output]})
			delete(a, output)
			delete(b, output)
		}
	}
	for output := range a {
		beforeOnly = append(beforeOnly, &actionOutputPair{output: output, before: before[output]})
	}
	for output := range b {
		beforeOnly = append(beforeOnly, &actionOutputPair{output: output, after: after[output]})
	}
	sort.Sort(beforeOnly)
	sort.Sort(afterOnly)
	sort.Sort(both)
	return
}
