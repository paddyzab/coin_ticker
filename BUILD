load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_prefix")

package(
		default_visibility = ["//visibility:public"],
)

go_binary(
		name = "coin_ticker",
		srcs = ["ct.go",
						"coin_ticker.go"],
		deps = ["@sling//:go_default_library"]
)

go_prefix("paddyzab/coin_ticker")
