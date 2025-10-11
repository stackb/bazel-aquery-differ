def _build_flags(ctx):
    """Helper to build common CLI flags."""
    flags = []
    if ctx.attr.unidiff:
        flags.append("--unidiff")
    if ctx.attr.cmpdiff:
        flags.append("--cmpdiff")
    if ctx.attr.serve:
        flags.append("--serve")
    if ctx.attr.open:
        flags.append("--open")
    return " ".join(flags)

_COMMON_TOOL_ATTRS = {
    "match": attr.string(
        doc = "strategy to compare before and after actions",
        default = "output_files",
        values = ["output_files", "mnemonic"],
    ),
    "serve": attr.bool(
        doc = "start webserver",
        default = True,
    ),
    "open": attr.bool(
        doc = "open browser to webserver URL",
        default = True,
    ),
    "unidiff": attr.bool(
        doc = "whether to compute unidiffs (can be slow)",
        default = False,
    ),
    "cmpdiff": attr.bool(
        doc = "whether to compute go-cmp diff (typically fast)",
        default = True,
    ),
    "_tool": attr.label(
        doc = "the aquerydiff tool",
        cfg = "exec",
        default = "//cmd/aquerydiff",
        executable = True,
    ),
}

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
    exit 1
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
{bazel} aquery --output=proto '{target}' > "$tmpdir/before.pb"

# checkout after commit and run aquery  
echo "Checking out after commit: {after_commit}"
git checkout {after_commit}
{bazel} aquery --output=proto '{target}' > "$tmpdir/after.pb"

# restore original commit
echo "Restoring original commit: $original_commit"
git checkout $original_commit

# run the tool
"$cwd/{tool}" {flags} --match '{match}' --before "$tmpdir/before.pb" --after "$tmpdir/after.pb" --report_dir=$cwd

# cleanup temporary directory
rm -rf "$tmpdir"
""".format(
            bazel = ctx.attr.bazel,
            tool = ctx.executable._tool.short_path,
            match = ctx.attr.match,
            target = ctx.attr.target,
            before_commit = ctx.attr.before,
            after_commit = ctx.attr.after,
            flags = _build_flags(ctx),
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
    attrs = dict(_COMMON_TOOL_ATTRS, **{
        "target": attr.string(
            doc = "bazel target to aquery",
            mandatory = True,
        ),
        "bazel": attr.string(
            doc = "the bazel executable",
            default = "bazel",
        ),
        "before": attr.string(
            doc = "the baseline git commit",
            mandatory = True,
        ),
        "after": attr.string(
            doc = "the after git commit",
            mandatory = True,
        ),
    }),
    executable = True,
)

def _aquery_diff_impl(ctx):
    files = [ctx.outputs.executable]
    runfiles = [ctx.executable._tool, ctx.file.before, ctx.file.after]

    ctx.actions.write(
        ctx.outputs.executable,
        """#!/bin/bash
set -euox pipefail

# switch to WORKSPACE
cwd=$(PWD)
before="$cwd/{before}"
after="$cwd/{after}"

cd $BUILD_WORKING_DIRECTORY

# run the tool
"$cwd/{tool}" {flags} --match '{match}' --before "$before" --after "$after" --report_dir=$cwd

""".format(
            tool = ctx.executable._tool.short_path,
            match = ctx.attr.match,
            before = ctx.file.before.short_path,
            after = ctx.file.after.short_path,
            flags = _build_flags(ctx),
        ),
        is_executable = True,
    )

    return [
        DefaultInfo(
            files = depset(files),
            runfiles = ctx.runfiles(files = runfiles),
        ),
    ]

aquery_diff = rule(
    implementation = _aquery_diff_impl,
    attrs = dict(_COMMON_TOOL_ATTRS, **{
        "before": attr.label(
            doc = "the baseline aquery file (proto, textproto, or jsonproto format)",
            allow_single_file = [".pb", ".proto", ".textproto", ".json", ".jsonproto"],
            mandatory = True,
        ),
        "after": attr.label(
            doc = "the comparison aquery file (proto, textproto, or jsonproto format)",
            allow_single_file = [".pb", ".proto", ".textproto", ".json", ".jsonproto"],
            mandatory = True,
        ),
    }),
    executable = True,
)
