"""
Test for generating rules from aquerydiff.
"""

load("@io_bazel_rules_go//go:def.bzl", "go_test")

def aquerydiff_generation_test(name, aquerydiff_binary, test_data, build_in_suffix = ".in", build_out_suffix = ".out", timeout_seconds = 2, size = None):
    """
    aquerydiff_generation_test is a macro for testing aquerydiff against workspaces.

    The generation test expects a file structure like the following:

    ```
    |-- <testDataPath>
        |-- some_test
            |-- WORKSPACE
            |-- README.md --> README describing what the test does.
            |-- expectedStdout.txt --> Expected stdout for this test.
            |-- expectedStderr.txt --> Expected stderr for this test.
            |-- expectedExitCode.txt --> Expected exit code for this test.
            |-- app
                |-- sourceFile.foo
                |-- BUILD.in --> BUILD file prior to running aquerydiff.
                |-- BUILD.out --> BUILD file expected after running aquerydiff.
    ```

    To update the expected files, run `UPDATE_SNAPSHOTS=true bazel run //path/to:the_test_target`.

    Args:
        name: The name of the test.
        aquerydiff_binary: The name of the aquerydiff binary target. For example, //path/to:my_aquerydiff.
        test_data: A list of target of the test data files you will pass to the test.
            This can be a https://bazel.build/reference/be/general#filegroup.
        build_in_suffix: The suffix for the input BUILD.bazel files. Defaults to .in.
            By default, will use files named BUILD.in as the BUILD files before running aquerydiff.
        build_out_suffix: The suffix for the expected BUILD.bazel files after running aquerydiff. Defaults to .out.
            By default, will use files named check the results of the aquerydiff run against files named BUILD.out.
        timeout_seconds: Number of seconds to allow the aquerydiff process to run before killing.
        size: Specifies a test target's "heaviness": how much time/resources it needs to run.
    """
    go_test(
        name = name,
        srcs = [Label("//pkg/generationtest:generation_test.go")],
        deps = [
            Label("//pkg/testtools"),
            "@io_bazel_rules_go//go/tools/bazel:go_default_library",
        ],
        args = [
            "-aquerydiff_binary_path=$(rootpath %s)" % aquerydiff_binary,
            "-build_in_suffix=%s" % build_in_suffix,
            "-build_out_suffix=%s" % build_out_suffix,
            "-timeout=%ds" % timeout_seconds,
        ],
        size = size,
        data = test_data + [
            aquerydiff_binary,
        ],
    )
