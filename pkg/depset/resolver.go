package depset

import (
	"fmt"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	"github.com/stackb/bazel-aquery-differ/pkg/artifact"
)

type Resolver struct {
	depSets         Map
	depSetArtifacts map[uint32][]string
	artifacts       artifact.PathMap
}

func NewResolver(artifacts artifact.PathMap, depSets Map) *Resolver {
	return &Resolver{
		artifacts:       artifacts,
		depSets:         depSets,
		depSetArtifacts: make(map[uint32][]string),
	}
}

func (r *Resolver) ResolveIds(depSetIds []uint32) (artifacts []string, err error) {
	for _, id := range depSetIds {
		depSet, ok := r.depSets[id]
		if !ok {
			return nil, fmt.Errorf("depSetOfFiles not found: %d", id)
		}
		files, err := r.Resolve(depSet)
		if err != nil {
			return nil, err
		}
		artifacts = append(artifacts, files...)
	}
	return
}

func (r *Resolver) Resolve(in *anpb.DepSetOfFiles) ([]string, error) {
	if depSetArtifacts, ok := r.depSetArtifacts[in.Id]; ok {
		return depSetArtifacts, nil
	}

	var artifacts []string

	for _, id := range in.DirectArtifactIds {
		artifact, ok := r.artifacts[id]
		if !ok {
			return nil, fmt.Errorf("artifact not found: %d", id)
		}
		artifacts = append(artifacts, artifact)
	}

	for _, id := range in.TransitiveDepSetIds {
		depSet, ok := r.depSets[id]
		if !ok {
			return nil, fmt.Errorf("depSetOfFiles not found: %d", id)
		}
		files, err := r.Resolve(depSet)
		if err != nil {
			return nil, fmt.Errorf("resolving transitive depSet %d: %w", id, err)
		}
		artifacts = append(artifacts, files...)
	}

	r.depSetArtifacts[in.Id] = artifacts

	return artifacts, nil
}
