load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "clientset.go",
        "doc.go",
    ],
    importpath = "github.com/mattmoor/warm-image/pkg/client/clientset/versioned",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/client/clientset/versioned/typed/warmimage/v2:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/k8s.io/client-go/discovery:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/client-go/util/flowcontrol:go_default_library",
    ],
)
