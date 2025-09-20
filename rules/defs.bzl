def _aquery_git_diff_impl(ctx):
    files = [ctx.outputs.executable]
    runfiles = [ctx.executable._tool]

    ctx.actions.write(
        ctx.outputs.executable,
        """#!/bin/bash
set -euox pipefail

# switch to WORKSPACE
cwd=$(PWD)
cd $BUILD_WORKING_DIRECTORY

# ensure no uncommitted changes
git diff --exit-code --quiet && git diff --cached --exit-code --quiet || {{
    echo "Error: Uncommitted changes detected"
    # exit 1
}}

# create temporary directory for aquery files
tmpdir="${{TMPDIR:-/tmp}}/aquery_diff_$$"
mkdir -p "$tmpdir"
echo "Using temporary directory: $tmpdir"

# store current git commit
original_commit=$(git rev-parse HEAD)
echo "Original commit: $original_commit"

# checkout before commit and run aquery
echo "Checking out before commit: {before_commit}"
git checkout {before_commit}
{bazel} aquery --output=proto {targets} > "$tmpdir/before.pb"

# checkout after commit and run aquery  
echo "Checking out after commit: {after_commit}"
git checkout {after_commit}
{bazel} aquery --output=proto {targets} > "$tmpdir/after.pb"

# restore original commit
echo "Restoring original commit: $original_commit"
git checkout $original_commit

# run the tool
"$cwd/{tool}" --before "$tmpdir/before.pb" --after "$tmpdir/after.pb" --report_dir=$cwd {serve_flag} {open_flag}

# cleanup temporary directory
rm -rf "$tmpdir"
""".format(
            bazel = ctx.attr.bazel,
            tool = ctx.executable._tool.short_path,
            targets = " ".join(ctx.attr.targets),
            before_commit = ctx.attr.before,
            after_commit = ctx.attr.after,
            serve_flag = "--serve" if ctx.attr.serve else "",
            open_flag = "--open" if ctx.attr.open else "",
        ),
        is_executable = True,
    )

    return [
        DefaultInfo(
            files = depset(files),
            runfiles = ctx.runfiles(files = runfiles),
        ),
    ]

aquery_git_diff = rule(
    implementation = _aquery_git_diff_impl,
    attrs = {
        "before": attr.string(
            doc = "the baseline git commit",
        ),
        "after": attr.string(
            doc = "the after git commit",
        ),
        "targets": attr.string_list(
            doc = "list of targets to build",
            mandatory = True,
        ),
        "bazel": attr.string(
            doc = "the bazel executable",
            default = "bazel",
        ),
        "serve": attr.bool(
            doc = "start webserver",
            default = True,
        ),
        "open": attr.bool(
            doc = "open browser to webserver URL",
            default = True,
        ),
        "_tool": attr.label(
            doc = "the aquerydiff tool",
            cfg = "exec",
            default = "//cmd/aquerydiff",
            executable = True,
        ),
    },
    executable = True,
)
