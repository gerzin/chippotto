load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "interpreter",
    srcs = [
        "assembler.go",
        "disassembler.go",
        "interpreter.go",
    ],
    importpath = "chippotto/interpreter",
    visibility = ["//visibility:public"],
)

go_test(
    name = "interpreter_test",
    srcs = ["interpreter_test.go"],
)
