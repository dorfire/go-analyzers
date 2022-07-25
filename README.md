# go-analyzers
Static analyzers for Go to enforce code consistency and discourage bad practices.

## Available analyzers

### visibilityorder

Breaking code example:

![image](https://user-images.githubusercontent.com/1499307/171998356-2f976920-cab5-48b7-8b9a-ec03eb7d5ee5.png)

Resulting compilation error:

![image](https://user-images.githubusercontent.com/1499307/171998390-9e413a54-c84b-4379-b00b-854f5be72a7b.png)

### onlyany

Forces the use of `any` (introduced in Go 1.8) instead of `interface{}`.

## Installation (with Bazel)

In your `WORKSPACE` file:
```starlark
http_archive(
    name = "com_github_dorfire_go_analyzers",
    sha256 = "2c07d07f67d7c402073548e8b1b36ffb23efa410d6958bacc4c12b02fcb1eac9",
    strip_prefix = "go-analyzers-master",
    urls = ["https://github.com/dorfire/go-analyzers/archive/refs/heads/master.tar.gz"],
)
```

In your NoGo target file (e.g. `build/BUILD.bazel`):
```starlark
load("@io_bazel_rules_go//go:def.bzl", "nogo")

nogo(
    name = "nogo",
    # ...
    deps = [
        "@com_github_dorfire_go_analyzers//src/onlyany",
        "@com_github_dorfire_go_analyzers//src/visibilityorder",
        # ...
    ],
)
```
