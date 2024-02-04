load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gui",
    srcs = ["renderer.go"],
    importpath = "chippotto/gui",
    visibility = ["//visibility:public"],
    deps = [
        "@io_fyne_fyne_v2//app",
        "@io_fyne_fyne_v2//container",
        "@io_fyne_fyne_v2//widget",
    ],
)
