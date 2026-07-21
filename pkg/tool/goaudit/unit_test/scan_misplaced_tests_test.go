package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"testing"
)

func TestMisplacedTestsFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/alfa/parse_test.go", "package alfa\n")
	v.WriteString("pkg/alfa/parse.go", "package alfa\n")
	result := scan.MisplacedTests(v)
	assert.Integer(t, 1, len(result))
	assert.String(t, scan.MisplacedTestKey, result[0].Key)
	assert.String(t, "pkg/alfa/parse_test.go", result[0].Path)
}

func TestMisplacedTestsHomesPass(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/alfa/unit_test/parse_test.go", "package unit_test\n")
	v.WriteString(
		"pkg/alfa/integration_test/client/client_test.go",
		"package client\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/unit_test/option_test.go",
		"package unit_test\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/integration_test/worker/worker_test.go",
		"package worker\n",
	)
	v.WriteString(
		"internal/bravo/unit_test/bravo_test.go",
		"package unit_test\n",
	)
	v.WriteString("charlie/unit_test/charlie_test.go", "package unit_test\n")
	assert.Integer(t, 0, len(scan.MisplacedTests(v)))
}

func TestMisplacedTestsTestdataExempt(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/alfa/deep/testdata/src/example/fixture_test.go",
		"package example\n",
	)
	assert.Integer(t, 0, len(scan.MisplacedTests(v)))
}

func TestMisplacedTestsNestedHomeFlagged(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString(
		"pkg/alfa/subsystem/integration_test/client/client_test.go",
		"package client\n",
	)
	result := scan.MisplacedTests(v)
	assert.Integer(t, 1, len(result))
	assert.String(
		t,
		"pkg/alfa/subsystem/integration_test/client/client_test.go",
		result[0].Path,
	)
}

func TestMisplacedTestsSorted(t *testing.T) {
	v := virtual_file_system.New()
	v.WriteString("pkg/zulu/zulu_test.go", "package zulu\n")
	v.WriteString("pkg/alfa/alfa_test.go", "package alfa\n")
	result := scan.MisplacedTests(v)
	assert.Integer(t, 2, len(result))
	assert.String(t, "pkg/alfa/alfa_test.go", result[0].Path)
	assert.String(t, "pkg/zulu/zulu_test.go", result[1].Path)
}
