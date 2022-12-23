workspace(name = "build_stack_bazel_aquery_differ")

load("//:repositories.bzl", "repositories")

repositories()

# ----------------------------------------------------
# @rules_proto
# ----------------------------------------------------

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies")

rules_proto_dependencies()

# ----------------------------------------------------
# @io_bazel_rules_go
# ----------------------------------------------------

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

go_rules_dependencies()

go_register_toolchains(version = "1.18.2")

# ----------------------------------------------------
# @bazel_gazelle
# ----------------------------------------------------
# gazelle:repository_macro go_repositories.bzl%go_repositories

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# ----------------------------------------------------
# @build_stack_rules_proto
# ----------------------------------------------------

register_toolchains("@build_stack_rules_proto//toolchain:standard")

# defining go_repository for @org_golang_google_grpc this in the WORKSPACE
# It must occur before gazelle_protobuf_extension_go_deps() macro call.
# I don't understand this one...  Despite 'bazel query //external:org_golang_google_grpc --output build'
# saying it's still coming from go_repos.bzl, that does not appear to be the case.
# Without this override org_golang_google_grpc is falling back to 1.27.0.
go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable_global",
    importpath = "google.golang.org/grpc",
    sum = "h1:XT2/MFpuPFsEX2fWh3YQtHkZ+WYZFQRfaUgLZYj/p6A=",
    version = "v1.42.0",
)

# ----------------------------------------------------
# @build_stack_rules_proto
# ----------------------------------------------------

load("@build_stack_rules_proto//:go_deps.bzl", "gazelle_protobuf_extension_go_deps")

gazelle_protobuf_extension_go_deps()

load("@build_stack_rules_proto//deps:go_core_deps.bzl", "go_core_deps")

go_core_deps()

# ----------------------------------------------------
# external go dependencies
# ----------------------------------------------------

load("//:go_repositories.bzl", "go_repositories")

go_repositories()

# ----------------------------------------------------
# external protobuf dependencies
# ----------------------------------------------------

load("//:proto_repositories.bzl", "proto_repositories")

proto_repositories()
