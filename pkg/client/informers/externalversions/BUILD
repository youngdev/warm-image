load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "factory.go",
        "generic.go",
    ],
    importpath = "github.com/mattmoor/warm-image/pkg/client/informers/externalversions",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/apis/warmimage/v2:go_default_library",
        "//pkg/client/clientset/versioned:go_default_library",
        "//pkg/client/informers/externalversions/internalinterfaces:go_default_library",
        "//pkg/client/informers/externalversions/warmimage:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/client-go/tools/cache:go_default_library",
    ],
)
