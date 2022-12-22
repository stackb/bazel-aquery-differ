load("@build_stack_rules_proto//rules/proto:proto_repository.bzl", "proto_repository")

def proto_repositories():
    protoapis()
    googleapis()
    remoteapis()
    bazelapis()

def googleapis():
    proto_repository(
        name = "googleapis",
        build_directives = [
            "gazelle:exclude google/example/endpointsapis/v1",
            "gazelle:exclude google/cloud/recommendationengine/v1beta1",  # is this a bug?
            "gazelle:proto_language go enabled true",
        ],
        build_file_expunge = True,
        build_file_proto_mode = "file",
        cfgs = ["//:rules_proto_config.yaml"],
        imports = ["@protoapis//:imports.csv"],
        override_go_googleapis = True,
        strip_prefix = "googleapis-02710fa0ea5312d79d7fb986c9c9823fb41049a9",
        type = "zip",
        urls = ["https://codeload.github.com/googleapis/googleapis/zip/02710fa0ea5312d79d7fb986c9c9823fb41049a9"],
    )

def bazelapis():
    proto_repository(
        name = "bazelapis",
        build_directives = [
            "gazelle:exclude third_party",
            "gazelle:proto_language go enable true",
            "gazelle:prefix github.com/bazelbuild/bazelapis",
        ],
        build_file_expunge = True,
        build_file_proto_mode = "file",
        cfgs = ["//:rules_proto_config.yaml"],
        imports = [
            "@googleapis//:imports.csv",
            "@protoapis//:imports.csv",
            "@remoteapis//:imports.csv",
        ],
        strip_prefix = "bazel-c2c49f430f2a7c277d21828e146dd6960dc0fbd6",
        type = "zip",
        urls = ["https://codeload.github.com/bazelbuild/bazel/zip/c2c49f430f2a7c277d21828e146dd6960dc0fbd6"],
    )

def remoteapis():
    proto_repository(
        name = "remoteapis",
        build_directives = [
            "gazelle:exclude third_party",
            "gazelle:exclude build/bazel/remote/asset/v1",
            "gazelle:exclude build/bazel/remote/logstream/v1",
            "gazelle:proto_language go enable true",
            "gazelle:proto_plugin protoc-gen-go option Mbuild/bazel/remote/execution/v2/remote_execution.proto=github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2",
            "gazelle:proto_plugin protoc-gen-go option Mbuild/bazel/semver/semver.proto=github.com/bazelbuild/remote-apis/build/bazel/semver",
            "gazelle:proto_plugin protoc-gen-go-grpc option Mbuild/bazel/remote/execution/v2/remote_execution.proto=github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2",
            "gazelle:proto_plugin protoc-gen-go-grpc option Mbuild/bazel/semver/semver.proto=github.com/bazelbuild/remote-apis/build/bazel/semver",
        ],
        build_file_expunge = True,
        build_file_proto_mode = "file",
        cfgs = ["//:rules_proto_config.yaml"],
        imports = [
            "@googleapis//:imports.csv",
            "@protoapis//:imports.csv",
        ],
        strip_prefix = "remote-apis-636121a32fa7b9114311374e4786597d8e7a69f3",
        type = "zip",
        urls = ["https://codeload.github.com/bazelbuild/remote-apis/zip/636121a32fa7b9114311374e4786597d8e7a69f3"],
    )

def protoapis():
    proto_repository(
        name = "protoapis",
        build_directives = [
            "gazelle:exclude testdata",
            "gazelle:exclude google/protobuf/compiler/ruby",
            "gazelle:proto_language go enable true",
        ],
        build_file_expunge = True,
        build_file_proto_mode = "file",
        cfgs = ["//:rules_proto_config.yaml"],
        deleted_files = [
            "google/protobuf/unittest_custom_options.proto",
            "google/protobuf/map_lite_unittest.proto",
            "google/protobuf/map_proto2_unittest.proto",
            "google/protobuf/test_messages_proto2.proto",
            "google/protobuf/test_messages_proto3.proto",
        ],
        strip_prefix = "protobuf-9650e9fe8f737efcad485c2a8e6e696186ae3862/src",
        type = "zip",
        urls = ["https://codeload.github.com/protocolbuffers/protobuf/zip/9650e9fe8f737efcad485c2a8e6e696186ae3862"],
    )
