module github.com/stackb/bazel-aquery-differ

go 1.23.1

require (
	github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2 v0.0.0-00010101000000-000000000000
	github.com/bazelbuild/rules_go v0.57.0
	github.com/google/go-cmp v0.7.0
	github.com/hexops/gotextdiff v1.0.3
	google.golang.org/protobuf v1.36.8
)

require (
	github.com/bazelbuild/bazel/src/main/protobuf/stardoc_output v0.0.0-00010101000000-000000000000 // indirect
	github.com/bazelbuild/bazelapis/src/main/protobuf/build v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/bazelbuild/bazel/src/main/protobuf/stardoc_output => ./genproto/github.com/bazelbuild/bazel/src/main/protobuf/stardoc_output

replace github.com/bazelbuild/bazelapis/src/main/protobuf/build => ./genproto/github.com/bazelbuild/bazelapis/src/main/protobuf/build

replace github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2 => ./genproto/github.com/bazelbuild/bazelapis/src/main/protobuf/analysis_v2
