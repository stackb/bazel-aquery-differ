package action

import (
	"sort"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
	"github.com/stackb/bazel-aquery-differ/pkg/artifact"
	"github.com/stackb/bazel-aquery-differ/pkg/depset"
	"github.com/stackb/bazel-aquery-differ/pkg/target"
)

// Graph holds compiled data about the action graph container.
type Graph struct {
	Container      *anpb.ActionGraphContainer
	Artifacts      artifact.PathMap
	Targets        target.Map
	DepSetOfFiles  depset.Map
	DepSetResolver depset.Resolver
	Actions        []*dipb.Action
	OutputMap      OutputMap
}

func NewGraph(container *anpb.ActionGraphContainer) (*Graph, error) {
	paths, err := artifact.NewPathMap(container.Artifacts, container.PathFragments)
	if err != nil {
		return nil, err
	}

	targets := target.NewMap(container.Targets)
	depSetOfFiles := depset.NewMap(container.DepSetOfFiles)
	depSetResolver := depset.NewResolver(paths, depSetOfFiles)

	actions := make([]*dipb.Action, len(container.Actions))
	for i, a := range container.Actions {
		action, err := NewAction(a, paths, targets, *depSetResolver)
		if err != nil {
			return nil, err
		}
		actions[i] = action
	}

	return &Graph{
		Container:      container,
		Artifacts:      paths,
		Targets:        targets,
		DepSetOfFiles:  depSetOfFiles,
		DepSetResolver: *depSetResolver,
		Actions:        actions,
		OutputMap:      NewOutputMap(actions),
	}, nil
}

func Partition(before, after OutputMap) (beforeOnly, afterOnly, both OutputPairs) {
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
			both = append(both, &OutputPair{Output: output, Before: before[output], After: after[output]})
			delete(a, output)
			delete(b, output)
		}
	}
	for output := range a {
		beforeOnly = append(beforeOnly, &OutputPair{Output: output, Before: before[output]})
	}
	for output := range b {
		beforeOnly = append(beforeOnly, &OutputPair{Output: output, After: after[output]})
	}
	sort.Sort(beforeOnly)
	sort.Sort(afterOnly)
	sort.Sort(both)
	return
}
