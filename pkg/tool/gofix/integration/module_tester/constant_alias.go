package module_tester

import "testing"

func ConstantAlias(t *testing.T) string {
	t.Helper()

	return New(t, "testmodule").
		File(
			"pkg/alfa/constant/constant.go",
			"package constant\n\nconst Alfa = \"alfa\"\n",
		).
		File(
			"pkg/bravo/constant/constant.go",
			"package constant\n\nconst Bravo = \"bravo\"\n",
		).
		File(
			"pkg/constant/constant.go",
			"package constant\n\nconst Root = \"root\"\n",
		).
		File(
			"pkg/golf/constant/constant.go",
			"package constant\n\nconst Golf = \"golf\"\n",
		).
		File(
			"pkg/juliet/constant/constant.go",
			"package constant\n\nconst Juliet = \"juliet\"\n",
		).
		File(
			"pkg/kilo/vocabulary/vocabulary.go",
			"package vocabulary\n\nconst Word = \"word\"\n",
		).
		File(
			"pkg/system/constant/constant.go",
			"package constant\n\nconst System = \"system\"\n",
		).
		File(
			"pkg/system/linux/constant/constant.go",
			"package constant\n\nconst Linux = \"linux\"\n",
		).
		File(
			"pkg/alfa/own.go",
			"package alfa\n\nimport (\n\town \"testmodule/pkg/alfa/constant\"\n\t\"testmodule/pkg/bravo/constant\"\n)\n\nfunc Own() string {\n\treturn own.Alfa + constant.Bravo\n}\n",
		).
		File(
			"pkg/charlie/lone.go",
			"package charlie\n\nimport \"testmodule/pkg/bravo/constant\"\n\nfunc Lone() string {\n\treturn constant.Bravo\n}\n",
		).
		File(
			"pkg/delta/root.go",
			"package delta\n\nimport \"testmodule/pkg/constant\"\n\nfunc Root() string {\n\treturn constant.Root\n}\n",
		).
		File(
			"pkg/system/linux/debian/deeper.go",
			"package debian\n\nimport (\n\t\"testmodule/pkg/system/constant\"\n\tlinux \"testmodule/pkg/system/linux/constant\"\n)\n\nfunc Deeper() string {\n\treturn constant.System + linux.Linux\n}\n",
		).
		File(
			"pkg/juliet/local.go",
			"package juliet\n\nimport (\n\tbravo \"testmodule/pkg/bravo/constant\"\n\town \"testmodule/pkg/juliet/constant\"\n)\n\nfunc Local() string {\n\tconstant := own.Juliet\n\n\treturn constant + bravo.Bravo\n}\n",
		).
		File(
			"pkg/kilo/held.go",
			"package kilo\n\nimport (\n\t\"testmodule/pkg/bravo/constant\"\n\tbravoConstant \"testmodule/pkg/kilo/vocabulary\"\n)\n\nfunc Held() string {\n\treturn constant.Bravo + bravoConstant.Word\n}\n",
		).
		File(
			"pkg/golf/conforming.go",
			"package golf\n\nimport (\n\tbravoConstant \"testmodule/pkg/bravo/constant\"\n\t\"testmodule/pkg/golf/constant\"\n)\n\nfunc Conforming() string {\n\treturn bravoConstant.Bravo + constant.Golf\n}\n",
		).
		Directory()
}
