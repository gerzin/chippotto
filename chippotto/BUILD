load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "chippotto_lib",
    srcs = ["main.go"],
    importpath = "github.com/gerzin/chippotto/chippotto",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "chippotto",
    embed = [":chippotto_lib"],
    visibility = ["//visibility:public"],
)
