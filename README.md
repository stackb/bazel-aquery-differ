[![CI](https://github.com/stackb/bazel-aquery-differ/actions/workflows/ci.yaml/badge.svg)](https://github.com/stackb/bazel-aquery-differ/actions/workflows/ci.yaml)

# bazel-aquery-differ

A tool to compare Bazel action query outputs with an interactive HTML report.
This is a re-imagination of the [Bazel
aquery_differ.py](https://github.com/bazelbuild/bazel/blob/master/tools/aquery_differ/aquery_differ.py)
in Go with enhanced visualization features.

## Features

- Compare Bazel action graphs between builds or git commits
- Interactive HTML reports with GitHub-style diff visualization
- Syntax-highlighted diffs (unified diff and go-cmp formats)
- Built-in web server with auto-open browser support
- Native Bazel rules for integration into your build

## Installation

### As a Bazel Module

Add to your `MODULE.bazel`:

```starlark
bazel_dep(name = "bazel-aquery-differ", version = "0.0.0")
```

> **Note**: This module is not yet published to the Bazel Central Registry. For
> now, use an `archive_override` or `git_override` pointing to this repository.

### As a Standalone Binary

Download a release artifact, or build from source:

```bash
git clone https://github.com/stackb/bazel-aquery-differ.git
cd bazel-aquery-differ
bazel build //cmd/aquerydiff
```

## Usage

### Using Bazel Rules

Load the rules in your `BUILD.bazel` file:

```starlark
load("@bazel-aquery-differ//rules:defs.bzl", "aquery_diff", "aquery_git_diff")
```

#### Rule: `aquery_diff`

Compare two aquery output files:

```starlark
aquery_diff(
    name = "compare_actions",
    before = "before.pb",
    after = "after.pb",
)
```

**Attributes:**

| Attribute | Type     | Default          | Description                                                                  |
|-----------|----------|------------------|------------------------------------------------------------------------------|
| `before`  | `label`  | **required**     | Baseline aquery file (`.pb`, `.proto`, `.textproto`, `.json`, `.jsonproto`)  |
| `after`   | `label`  | **required**     | Comparison aquery file (same format options)                                 |
| `match`   | `string` | `"output_files"` | Strategy to match before and after actions: `"output_files"` or `"mnemonic"` |
| `serve`   | `bool`   | `True`           | Start web server to view report                                              |
| `open`    | `bool`   | `True`           | Automatically open browser to report                                         |
| `unidiff` | `bool`   | `False`          | Generate unified diffs (can be slow for large actions)                       |
| `cmpdiff` | `bool`   | `True`           | Generate go-cmp diffs (fast, structural comparison)                          |

Run the comparison:

```bash
bazel run //path/to:compare_actions
```

> **Performance Note**: The `unidiff` attribute defaults to `False` because generating unified diffs can be prohibitively slow for large actions with many inputs/outputs. The `cmpdiff` format (enabled by default) is much faster and provides good structural comparison for most use cases. Only enable `unidiff` if you need the traditional unified diff format and are willing to wait for the additional processing time.

**Choosing a Match Strategy:**

The `match` attribute determines how actions are paired between the before and
after builds:

- **`output_files`** (default): Actions are matched by their output file paths.
  Use this when comparing the same target across different commits or
  configurations. This is the most common use case and ensures you're comparing
  the exact same action that produces the same outputs.

- **`mnemonic`**: Actions are matched by their mnemonic (action type, e.g.,
  "GoCompile", "CppCompile"). Use this when comparing different targets that use
  similar build rules. For example, comparing `//old/pkg:binary` vs
  `//new/pkg:binary` where both are `go_binary` targets but produce different
  output paths. This helps identify how the same type of action differs between
  targets.

Example using mnemonic matching:

```starlark
aquery_diff(
    name = "compare_go_binaries",
    before = "old_binary.pb",
    after = "new_binary.pb",
    match = "mnemonic",  # Compare by action type instead of output path
)
```

#### Rule: `aquery_git_diff`

Compare aquery outputs between git commits:

```starlark
aquery_git_diff(
    name = "git_compare",
    before = "main",
    after = "feature-branch",
    target = "//my/package:target",
)
```

**Attributes:**

Same as `aquery_diff`, plus:

| Attribute | Type     | Default      | Description                                                |
|-----------|----------|--------------|------------------------------------------------------------|
| `target`  | `string` | **required** | Bazel target to aquery (e.g., `//pkg:binary`, `deps(...)`) |
| `bazel`   | `string` | `"bazel"`    | Path to bazel executable                                   |
| `before`  | `string` | **required** | Git commit/branch/tag for baseline                         |
| `after`   | `string` | **required** | Git commit/branch/tag for comparison                       |

This rule will:
1. Check for uncommitted changes (fails if found)
2. Checkout `before` commit and run `bazel aquery`
3. Checkout `after` commit and run `bazel aquery`
4. Restore original commit
5. Generate comparison report

### Using the CLI

Generate aquery files using Bazel:

```bash
# Binary proto format (recommended for large graphs)
bazel aquery //pkg:target --output=proto > before.pb

# Text proto format (human-readable)
bazel aquery //pkg:target --output=textproto > before.textproto

# JSON proto format
bazel aquery //pkg:target --output=jsonproto > before.json
```

> **Supported formats**: The tool automatically detects format based on file extension:
> - Binary: `.pb`, `.proto`
> - Text: `.textproto`
> - JSON: `.json`, `.jsonproto`

Run the comparison:

```bash
aquerydiff \
  --before before.pb \
  --after after.pb \
  --report_dir ./output \
  --serve \
  --open
```

**CLI Flags:**

- `--before` - Path to baseline aquery file
- `--after` - Path to comparison aquery file
- `--report_dir` - Directory to write HTML report
- `--match` - Matching strategy: `output_files` (default) or `mnemonic`
- `--serve` - Start web server (default: true)
- `--open` - Open browser automatically (default: true)
- `--unidiff` - Generate unified diffs (default: false)
- `--cmpdiff` - Generate go-cmp diffs (default: true)

> **Note**: The report title is automatically derived from the most common target in the action graph.

### Report Output

The HTML report shows:

- **Actions only in before** - Removed actions
- **Actions only in after** - New actions
- **Non-equal actions** - Actions with changes
- **Equal actions** - Unchanged actions

Each action displays:
- Mnemonic (action type)
- Output files
- Links to before/after JSON/textproto representations
- Colorized diffs (unified and/or go-cmp format)

<img width="934" alt="Example report showing action comparison" src="https://user-images.githubusercontent.com/50580/209453563-064db4dd-4068-4d2f-8bb3-35c425bfb8b5.png">

## Example

See the [examples/simple](examples/simple) directory for working examples using both rules.

## Contributing

Contributions welcome! Please open an issue or pull request.

## License

Apache 2.0
