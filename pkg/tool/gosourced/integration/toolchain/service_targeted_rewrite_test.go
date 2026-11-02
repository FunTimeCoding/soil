package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"strings"
	"testing"
)

func TestTargetedRenameFunctionMatchesFullLoad(t *testing.T) {
	tree := sameRewrite(
		t,
		func(s *service_tester.Services, d string) (*output.Results, error) {
			return s.Symbol.Rename(
				t.Context(),
				d,
				"example/alfa",
				"NewServer",
				"Make",
				"",
				false,
			)
		},
	)
	assert.Integer(t, 5, strings.Count(tree, "Make()"))
}

func TestTargetedRenameFieldMatchesFullLoad(t *testing.T) {
	tree := sameRewrite(
		t,
		func(s *service_tester.Services, d string) (*output.Results, error) {
			return s.Symbol.Rename(
				t.Context(),
				d,
				"example/alfa",
				"Port",
				"Number",
				"Server",
				false,
			)
		},
	)
	assert.Integer(t, 0, strings.Count(tree, "Port"))
}

func TestTargetedMovePackageMatchesFullLoad(t *testing.T) {
	tree := sameRewrite(
		t,
		func(s *service_tester.Services, d string) (*output.Results, error) {
			return s.Packages.MovePackage(
				t.Context(),
				d,
				"example/delta",
				"example/zulu/delta",
				false,
			)
		},
	)
	assert.True(t, strings.Contains(tree, "\"example/zulu/delta\""))
}
