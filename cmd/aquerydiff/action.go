package main

import (
	"fmt"
	"sort"
	"strings"

	anpb "github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2"
)

// Action is the same struct as anpb.Action but with identifiers resolved to the
// real names.
type Action struct {
	// The target that was responsible for the creation of the action.
	Target string

	// The aspects that were responsible for the creation of the action (if
	// any). In the case of aspect-on-aspect, AspectDescriptors are listed in
	// topological order of the dependency graph. e.g. [A, B] would imply that
	// aspect A is applied on top of aspect B.
	//   repeated uint32 aspect_descriptor_ids = 2;

	// Encodes all significant behavior that might affect the output. The key
	// must change if the work performed by the execution of this action
	// changes. Note that the key doesn't include checksums of the input files.
	ActionKey string

	// The mnemonic for this kind of action.
	Mnemonic string

	// The configuration under which this action is executed.
	//uint32 configuration_id = 5;

	// The command line arguments of the action. This will be only set if
	// explicitly requested.
	Arguments []string

	// The list of environment variables to be set before executing the command.
	EnvironmentVariables []*anpb.KeyValuePair

	// The set of input dep sets that the action depends upon. If the action
	// does input discovery, the contents of this set might change during
	// execution.
	Inputs []string

	// The list of Artifact IDs that represent the output files that this action
	// will generate.
	Outputs []string

	// a space-concatenated string of the .Outputs
	OutputFiles string

	// True iff the action does input discovery during execution.
	DiscoversInputs bool

	// Execution info for the action.  Remote execution services may use this
	// information to modify the execution environment, but actions will
	// generally not be aware of it.
	ExecutionInfo []*anpb.KeyValuePair

	// The list of param files. This will be only set if explicitly requested.
	ParamFiles []*anpb.ParamFile

	// The id to an Artifact that is the primary output of this action.
	PrimaryOutput string

	// The execution platform for this action. Empty if the action has no
	// execution platform.
	ExecutionPlatform string

	// The template content of the action, if it is TemplateExpand action.
	TemplateContent string

	// The list of substitution should be performed on the template. The key is
	// the string to be substituted and the value is the string to be substituted
	// to.
	Substitutions []*anpb.KeyValuePair

	// The contents of the file for the actions.write() action
	// (guarded by the --include_file_write_contents flag).
	FileContents string
}

func newAction(in *anpb.Action, artifacts artifactPathMap, targets targetMap, deps depSetResolver) (*Action, error) {
	target, ok := targets[in.TargetId]
	if !ok {
		return nil, fmt.Errorf("target not found: %d", in.TargetId)
	}
	inputs, err := deps.resolveIds(in.InputDepSetIds)
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
	out := &Action{
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
