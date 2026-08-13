package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestConstantPlacementStrayFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/alfa/parse.go", "package alfa\n")
	v.WriteString("pkg/alfa/constant.go", "package alfa\n")
	result := scan.ConstantPlacement(v)
	assert.Integer(t, 1, len(result))
	assert.String(t, "constant_file", result[0].Key)
	assert.String(t, "pkg/alfa/constant.go", result[0].Path)
}

func TestConstantPlacementHomesPass(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/constant/constant.go", "package constant\n")
	v.WriteString("pkg/alfa/constant/constant.go", "package constant\n")
	v.WriteString("pkg/alfa/constant/style.go", "package constant\n")
	v.WriteString("pkg/tool/gotestd/constant/constant.go", "package constant\n")
	assert.Integer(t, 0, len(scan.ConstantPlacement(v)))
}

func TestConstantPlacementDepthFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/alfa/deep/constant/constant.go", "package constant\n")
	v.WriteString("cmd/example/alfa/constant/constant.go", "package constant\n")
	v.WriteString("pkg/tool/constant/constant.go", "package constant\n")
	result := scan.ConstantPlacement(v)
	assert.Integer(t, 3, len(result))
	assert.String(t, "constant_depth", result[0].Key)
	assert.String(t, "cmd/example/alfa/constant", result[0].Path)
	assert.String(t, "pkg/alfa/deep/constant", result[1].Path)
	assert.String(t, "pkg/tool/constant", result[2].Path)
}

func TestConstantPlacementNestedFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/constant/emoji/emoji.go", "package emoji\n")
	result := scan.ConstantPlacement(v)
	assert.Integer(t, 1, len(result))
	assert.String(t, "constant_nested", result[0].Key)
	assert.String(t, "pkg/constant/emoji", result[0].Path)
}

func TestConstantPlacementNestedStrayPair(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/constant/directory/constant.go", "package directory\n")
	result := scan.ConstantPlacement(v)
	assert.Integer(t, 2, len(result))
	assert.String(t, "constant_nested", result[0].Key)
	assert.String(t, "pkg/constant/directory", result[0].Path)
	assert.String(t, "constant_file", result[1].Key)
	assert.String(t, "pkg/constant/directory/constant.go", result[1].Path)
}

func TestConstantPlacementTestHomePasses(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/constant/unit_test/constant_test.go",
		"package unit_test\n",
	)
	v.WriteString(
		"pkg/alfa/constant/integration_test/anchor_test.go",
		"package integration_test\n",
	)
	assert.Integer(t, 0, len(scan.ConstantPlacement(v)))
}

func TestConstantPlacementTestdataExempt(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/alfa/testdata/src/pkg/target/constant.go",
		"package target\n",
	)
	v.WriteString(
		"pkg/alfa/testdata/src/pkg/deep/constant/constant.go",
		"package constant\n",
	)
	assert.Integer(t, 0, len(scan.ConstantPlacement(v)))
}

func TestConstantPlacementSorted(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/zulu/constant.go", "package zulu\n")
	v.WriteString("pkg/alfa/constant.go", "package alfa\n")
	result := scan.ConstantPlacement(v)
	assert.Integer(t, 2, len(result))
	assert.String(t, "pkg/alfa/constant.go", result[0].Path)
	assert.String(t, "pkg/zulu/constant.go", result[1].Path)
}
