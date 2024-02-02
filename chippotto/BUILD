load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "chippotto_lib",
    srcs = ["main.go"],
    importpath = "chippotto",
    visibility = ["//visibility:private"],
    deps = ["//chippotto/interpreter"],
)

go_binary(
    name = "chippotto",
    embed = [":chippotto_lib"],
    visibility = ["//visibility:public"],
)
