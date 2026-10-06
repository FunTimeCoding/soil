package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service"
	"strings"
	"testing"
)

func TestTargetedRenameFunctionMatchesFullLoad(t *testing.T) {
	tree := sameRewrite(
		t,
		func(s *service.Service, d string) (*output.Results, error) {
			return s.Rename(d, "example/alfa", "NewServer", "Make", "", false)
		},
	)
	assert.Integer(t, 5, strings.Count(tree, "Make()"))
}

func TestTargetedRenameFieldMatchesFullLoad(t *testing.T) {
	tree := sameRewrite(
		t,
		func(s *service.Service, d string) (*output.Results, error) {
			return s.Rename(
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
		func(s *service.Service, d string) (*output.Results, error) {
			return s.MovePackage(
				d,
				"example/delta",
				"example/zulu/delta",
				false,
			)
		},
	)
	assert.True(t, strings.Contains(tree, "\"example/zulu/delta\""))
}
