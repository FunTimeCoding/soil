package module_tester

import "testing"

func CollisionKinds(t *testing.T) string {
	t.Helper()

	return New(t, "testmodule").
		File(
			"pkg/core/v1/v1.go",
			"package v1\n\ntype Pod struct{}\n",
		).
		File(
			"pkg/meta/v1/v1.go",
			"package v1\n\ntype Object struct{}\n",
		).
		File(
			"pkg/alfa/pod.go",
			"package alfa\n\nimport (\n\t\"testmodule/pkg/core/v1\"\n\tmeta \"testmodule/pkg/meta/v1\"\n)\n\nfunc Pod() (*v1.Pod, *meta.Object) {\n\treturn &v1.Pod{}, &meta.Object{}\n}\n",
		).
		File(
			"pkg/bravo/chain/chain.go",
			"package chain\n\ntype Chain struct{}\n",
		).
		File(
			"pkg/charlie/chain/chain.go",
			"package chain\n\ntype Link struct{}\n",
		).
		File(
			"pkg/bravo/lower/lower.go",
			"package lower\n\nimport (\n\tchain1 \"testmodule/pkg/bravo/chain\"\n\t\"testmodule/pkg/charlie/chain\"\n)\n\nfunc Lower(*chain.Link) *chain1.Chain {\n\treturn &chain1.Chain{}\n}\n",
		).
		File(
			"chain.go",
			"package chain\n\ntype Root struct{}\n",
		).
		File(
			"pkg/delta/chain/chain.go",
			"package chain\n\ntype Delta struct{}\n",
		).
		File(
			"pkg/delta/held/held.go",
			"package held\n\nimport (\n\t\"testmodule\"\n\tchain1 \"testmodule/pkg/delta/chain\"\n)\n\nfunc Held(*chain.Root) *chain1.Delta {\n\treturn &chain1.Delta{}\n}\n",
		).
		File(
			"pkg/errors/errors.go",
			"package errors\n\nfunc Echo() string { return \"echo\" }\n",
		).
		File(
			"pkg/echo/fail.go",
			"package echo\n\nimport (\n\t\"errors\"\n\tlibrary \"testmodule/pkg/errors\"\n)\n\nfunc Fail() error {\n\treturn errors.New(library.Echo())\n}\n",
		).
		File(
			"pkg/face/face.go",
			"package face\n\ntype Reporter interface{}\n",
		).
		File(
			"pkg/golf/face/face.go",
			"package face\n\ntype Client interface{}\n",
		).
		File(
			"pkg/golf/service/service.go",
			"package service\n\nimport (\n\t\"testmodule/pkg/face\"\n\tgolfFace \"testmodule/pkg/golf/face\"\n)\n\nfunc Service(face.Reporter, golfFace.Client) {}\n",
		).
		File(
			"pkg/hotel/mock_client/client.go",
			"package mock_client\n\ntype Client struct{}\n",
		).
		File(
			"pkg/tool/goindia/client/mock_client/client.go",
			"package mock_client\n\ntype Client struct{}\n",
		).
		File(
			"pkg/tool/goindia/unit/doubles.go",
			"package unit\n\nimport (\n\t\"testmodule/pkg/hotel/mock_client\"\n\tindiaMock \"testmodule/pkg/tool/goindia/client/mock_client\"\n)\n\nfunc Doubles() (*mock_client.Client, *indiaMock.Client) {\n\treturn &mock_client.Client{}, &indiaMock.Client{}\n}\n",
		).
		Directory()
}
