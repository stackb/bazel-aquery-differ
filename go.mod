module github.com/stackb/bazel-aquery-differ

go 1.23.1

require (
	github.com/bazelbuild/bazelapis/src/main/java/com/google/devtools/build/lib/buildeventstream/proto/build_event_stream/build_event_stream v0.0.0-00010101000000-000000000000
	github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2 v0.0.0-00010101000000-000000000000
	github.com/bazelbuild/rules_go v0.57.0
	github.com/google/go-cmp v0.7.0
	github.com/hexops/gotextdiff v1.0.3
	google.golang.org/protobuf v1.36.8
)

require (
	github.com/bazelbuild/bazel/src/main/java/com/google/devtools/build/lib/packages/metrics v0.0.0-00010101000000-000000000000 // indirect
	github.com/bazelbuild/bazel/src/main/protobuf/stardoc_output v0.0.0-00010101000000-000000000000 // indirect
	github.com/bazelbuild/bazel/src/main/protobuf/strategy_policy v0.0.0-00010101000000-000000000000 // indirect
	github.com/bazelbuild/bazelapis/src/main/protobuf/action_cache v0.0.0-00010101000000-000000000000 // indirect
	github.com/bazelbuild/bazelapis/src/main/protobuf/build v0.0.0-00010101000000-000000000000 // indirect
	github.com/bazelbuild/bazelapis/src/main/protobuf/command_line v0.0.0-00010101000000-000000000000 // indirect
	github.com/bazelbuild/bazelapis/src/main/protobuf/failure_details v0.0.0-00010101000000-000000000000 // indirect
	github.com/bazelbuild/bazelapis/src/main/protobuf/invocation_policy v0.0.0-00010101000000-000000000000 // indirect
	github.com/bazelbuild/bazelapis/src/main/protobuf/option_filters v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/bazelbuild/bazelapis/src/main/protobuf/build => ./genproto/github.com/bazelbuild/bazelapis/src/main/protobuf/build

replace github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2 => ./genproto/github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2

replace github.com/bazelbuild/bazel/src/main/protobuf/stardoc_output => ./genproto/github.com/bazelbuild/bazel/src/main/protobuf/stardoc_output

replace google.golang.org/protobuf/types/known/durationpb => ./genproto/google.golang.org/protobuf/types/known/durationpb

replace github.com/bazelbuild/bazel/src/main/java/com/google/devtools/build/lib/packages/metrics => ./genproto/github.com/bazelbuild/bazel/src/main/java/com/google/devtools/build/lib/packages/metrics

replace github.com/bazelbuild/bazelapis/src/main/protobuf/action_cache => ./genproto/github.com/bazelbuild/bazelapis/src/main/protobuf/action_cache

replace github.com/bazelbuild/bazelapis/src/main/protobuf/option_filters => ./genproto/github.com/bazelbuild/bazelapis/src/main/protobuf/option_filters

replace github.com/bazelbuild/bazelapis/src/main/protobuf/command_line => ./genproto/github.com/bazelbuild/bazelapis/src/main/protobuf/command_line

replace google.golang.org/protobuf/types/descriptorpb => ./genproto/google.golang.org/protobuf/types/descriptorpb

replace github.com/bazelbuild/bazelapis/src/main/protobuf/failure_details => ./genproto/github.com/bazelbuild/bazelapis/src/main/protobuf/failure_details

replace github.com/bazelbuild/bazel/src/main/protobuf/strategy_policy => ./genproto/github.com/bazelbuild/bazel/src/main/protobuf/strategy_policy

replace github.com/bazelbuild/bazelapis/src/main/protobuf/invocation_policy => ./genproto/github.com/bazelbuild/bazelapis/src/main/protobuf/invocation_policy

replace google.golang.org/protobuf/types/known/anypb => ./genproto/google.golang.org/protobuf/types/known/anypb

replace google.golang.org/protobuf/types/known/timestamppb => ./genproto/google.golang.org/protobuf/types/known/timestamppb

replace github.com/bazelbuild/bazelapis/src/main/java/com/google/devtools/build/lib/buildeventstream/proto/build_event_stream/build_event_stream => ./genproto/github.com/bazelbuild/bazelapis/src/main/java/com/google/devtools/build/lib/buildeventstream/proto/build_event_stream/build_event_stream
