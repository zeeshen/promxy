load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "common.go",
        "client.go",
        "server.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//env:go_default_library",
        "//defines:go_default_library",
        "@com_github_golang_glog//:go_default_library",
        "@io_opencensus_go//trace:go_default_library",
        "@io_opencensus_go//plugin/ochttp:go_default_library",
    ],
)
