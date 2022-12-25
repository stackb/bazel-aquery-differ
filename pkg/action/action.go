package action

import (
	"fmt"
	"sort"
	"strings"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
	dipb "github.com/stackb/bazel-aquery-differ/build/stack/bazel/aquery/differ"
	"github.com/stackb/bazel-aquery-differ/pkg/artifact"
	"github.com/stackb/bazel-aquery-differ/pkg/depset"
	"github.com/stackb/bazel-aquery-differ/pkg/target"
)

func NewAction(id string, in *anpb.Action, artifacts artifact.PathMap, targets target.Map, deps depset.Resolver) (*dipb.Action, error) {
	target, ok := targets[in.TargetId]
	if !ok {
		return nil, fmt.Errorf("target not found: %d", in.TargetId)
	}
	inputs, err := deps.ResolveIds(in.InputDepSetIds)
	if err != nil {
		return nil, err
	}
	sort.Strings(inputs)
	outputs := make([]string, len(in.OutputIds))
	for i, id := range in.OutputIds {
		artifact, ok := artifacts[id]
		if !ok {
			return nil, fmt.Errorf("output artifact not found: %d", id)
		}
		outputs[i] = artifact
	}
	sort.Strings(outputs)
	out := &dipb.Action{
		Id:                   id,
		Target:               target.Label,
		ActionKey:            in.ActionKey,
		Mnemonic:             in.Mnemonic,
		Arguments:            in.Arguments,
		EnvironmentVariables: in.EnvironmentVariables,
		Inputs:               inputs,
		Outputs:              outputs,
		OutputFiles:          strings.Join(outputs, " "),
		DiscoversInputs:      in.DiscoversInputs,
		ExecutionInfo:        in.ExecutionInfo,
		ParamFiles:           in.ParamFiles,
		PrimaryOutput:        ">",
		ExecutionPlatform:    in.ExecutionPlatform,
		TemplateContent:      in.TemplateContent,
		Substitutions:        in.Substitutions,
		// FileContents:         in.FileContent,
	}

	return out, nil
}
