platforms = [
    struct(
        arch = "amd64",
        gc_linkopts = [
            "-s",
            "-w",
        ],
        os = "darwin",
    ),
    struct(
        arch = "arm64",
        gc_linkopts = [
            "-s",
            "-w",
        ],
        os = "darwin",
    ),
    struct(
        arch = "amd64",
        gc_linkopts = [
            "-s",
            "-w",
        ],
        os = "linux",
    ),
    struct(
        arch = "arm64",
        gc_linkopts = [
            "-s",
            "-w",
        ],
        os = "linux",
    ),
    struct(
        arch = "amd64",
        gc_linkopts = [],
        os = "windows",
    ),
]
