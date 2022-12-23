package main

import (
	"fmt"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
)

// artifactPathMap maps artifact.Id to the output path of the artifact.
type artifactPathMap map[uint32]string

// newArtifactPathMap creates a new artifactPathMap based on the given list of
// artifacts and the path fragments.
func newArtifactPathMap(artifacts []*anpb.Artifact, pathFragments []*anpb.PathFragment) (artifactPathMap, error) {
	pathResolver := newPathFragmentResolver(pathFragments)
	result := make(artifactPathMap)
	for _, artifact := range artifacts {
		path, err := pathResolver.resolve(artifact.PathFragmentId)
		if err != nil {
			return nil, fmt.Errorf("resolving path for artifact %v: %w", artifact, err)
		}
		result[artifact.Id] = path
	}
	return result, nil
}
