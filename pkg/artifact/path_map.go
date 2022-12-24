package artifact

import (
	"fmt"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	"github.com/stackb/bazel-aquery-differ/pkg/pathfragment"
)

// PathMap maps artifact.Id to the output path of the artifact.
type PathMap map[uint32]string

// NewPathMap creates a new artifactPathMap based on the given list of artifacts
// and the path fragments.
func NewPathMap(artifacts []*anpb.Artifact, pathFragments []*anpb.PathFragment) (PathMap, error) {
	pathResolver := pathfragment.NewResolver(pathFragments)
	result := make(PathMap)
	for _, artifact := range artifacts {
		path, err := pathResolver.Resolve(artifact.PathFragmentId)
		if err != nil {
			return nil, fmt.Errorf("resolving path for artifact %v: %w", artifact, err)
		}
		result[artifact.Id] = path
	}
	return result, nil
}
