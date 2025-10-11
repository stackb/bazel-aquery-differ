package action

import (
	"fmt"
	"sort"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	dipb "github.com/stackb/bazel_difftools/build/stack/bazel/aquery/differ"
	"github.com/stackb/bazel_difftools/pkg/artifact"
	"github.com/stackb/bazel_difftools/pkg/depset"
	"github.com/stackb/bazel_difftools/pkg/target"
)

// Graph holds compiled data about the action graph container.
type Graph struct {
	Name           string
	Container      *anpb.ActionGraphContainer
	Artifacts      artifact.PathMap
	Targets        target.Map
	DepSetOfFiles  depset.Map
	DepSetResolver depset.Resolver
	Actions        []*dipb.Action
}

func NewGraph(name string, container *anpb.ActionGraphContainer) (*Graph, error) {
	paths, err := artifact.NewPathMap(container.Artifacts, container.PathFragments)
	if err != nil {
		return nil, err
	}

	targets := target.NewMap(container.Targets)
	depSetOfFiles := depset.NewMap(container.DepSetOfFiles)
	depSetResolver := depset.NewResolver(paths, depSetOfFiles)

	actions := make([]*dipb.Action, len(container.Actions))
	for i, a := range container.Actions {
		id := fmt.Sprintf("%s/%d", name, i)
		action, err := NewAction(id, a, paths, targets, *depSetResolver)
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
	}, nil
}

// GetPrimaryTarget returns the most common target from the actions in the graph,
// or empty string if there are no actions.
func (g *Graph) GetPrimaryTarget() string {
	if len(g.Actions) == 0 {
		return ""
	}

	// Count target occurrences
	targetCounts := make(map[string]int)
	for _, action := range g.Actions {
		if action.Target != "" {
			targetCounts[action.Target]++
		}
	}

	// Find the most common target
	var maxTarget string
	var maxCount int
	for target, count := range targetCounts {
		if count > maxCount {
			maxTarget = target
			maxCount = count
		}
	}

	return maxTarget
}

func Partition(before, after ActionMap) (beforeOnly, afterOnly, both OutputPairs) {
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
			both = append(both, &OutputPair{
				Output: output,
				Action: before[output],
				Before: before[output],
				After:  after[output],
			})
			delete(a, output)
			delete(b, output)
		}
	}
	for output := range a {
		beforeOnly = append(beforeOnly, &OutputPair{
			Output: output,
			Action: before[output],
			Before: before[output],
		})
	}
	for output := range b {
		beforeOnly = append(beforeOnly, &OutputPair{
			Output: output,
			Action: after[output],
			After:  after[output],
		})
	}
	sort.Sort(beforeOnly)
	sort.Sort(afterOnly)
	sort.Sort(both)
	return
}
