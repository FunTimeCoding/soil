package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestStrayHelpersFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/alfa/unit/parse_test.go", "package unit\n")
	v.WriteString("pkg/alfa/unit/fixture.go", "package unit\n")
	result := scan.StrayHelpers(v)
	assert.Integer(t, 1, len(result))
	assert.String(t, "stray_helper", result[0].Key)
	assert.String(t, "pkg/alfa/unit/fixture.go", result[0].Path)
}

func TestStrayHelpersPackagePasses(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/alfa/unit/parse_test.go", "package unit\n")
	v.WriteString(
		"pkg/alfa/unit/parse_tester/fixture.go",
		"package parse_tester\n",
	)
	assert.Integer(t, 0, len(scan.StrayHelpers(v)))
}

func TestStrayHelpersWithoutTestsPass(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/alfa/parse.go", "package alfa\n")
	v.WriteString("pkg/alfa/parse_test.go", "package alfa\n")
	assert.Integer(t, 0, len(scan.StrayHelpers(v)))
}

func TestStrayHelpersTestdataExempt(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/alfa/unit/testdata/src/example/fixture_test.go",
		"package example\n",
	)
	v.WriteString(
		"pkg/alfa/unit/testdata/src/example/fixture.go",
		"package example\n",
	)
	assert.Integer(t, 0, len(scan.StrayHelpers(v)))
}

func TestStrayHelpersIntegrationSubdirectoryFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/tool/alfad/integration/store/store_test.go",
		"package store\n",
	)
	v.WriteString("pkg/tool/alfad/integration/store/seed.go", "package store\n")
	result := scan.StrayHelpers(v)
	assert.Integer(t, 1, len(result))
	assert.String(t, "pkg/tool/alfad/integration/store/seed.go", result[0].Path)
}

func TestStrayHelpersSorted(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/zulu/unit/zulu_test.go", "package unit\n")
	v.WriteString("pkg/zulu/unit/zulu.go", "package unit\n")
	v.WriteString("pkg/alfa/unit/alfa_test.go", "package unit\n")
	v.WriteString("pkg/alfa/unit/alfa.go", "package unit\n")
	result := scan.StrayHelpers(v)
	assert.Integer(t, 2, len(result))
	assert.String(t, "pkg/alfa/unit/alfa.go", result[0].Path)
	assert.String(t, "pkg/zulu/unit/zulu.go", result[1].Path)
}
