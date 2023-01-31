[![CI](https://github.com/stackb/bazel-aquery-differ/actions/workflows/ci.yaml/badge.svg)](https://github.com/stackb/bazel-aquery-differ/actions/workflows/ci.yaml)

# bazel-aquery-differ

This is a port of
<https://github.com/bazelbuild/bazel/blob/master/tools/aquery_differ/aquery_differ.py>
to golang.

## Installation

Download and unzip a release artifact, or clone and `bazel build //cmd/aquerydiff`.

## Usage

```bash
aquerydiff --before <BEFORE_FILE> --after <AFTER_FILE> --report_dir <REPORT_DIR>
```

You can generate the `<BEFORE_FILE>` (and `<AFTER_FILE>`) using:

```bash
bazel aquery //pkg:target-name --output jsonproto > before.json
bazel aquery //pkg:target-name --output textproto > before.textproto
bazel aquery //pkg:target-name --output proto > before.pb
```

> The file extensions are relevant; the proto decoder will be `protojson` if
`.json`, `prototext` if `.textproto` and `proto` otherwise.


An HTML report and accessory files will be written to the given `--report_dir`,
which you could serve as follows:

```
(cd <REPORT_DIR> && python3 -m http.server 8000) &
```

> Report will look something like:

<img width="934" alt="image" src="https://user-images.githubusercontent.com/50580/209453563-064db4dd-4068-4d2f-8bb3-35c425bfb8b5.png">
