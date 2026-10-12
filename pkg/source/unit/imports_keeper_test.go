package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"testing"
)

func TestKeeperIsTheStandardLibrary(t *testing.T) {
	keeper, reason := imports.Keeper(
		"example.org/alfa/pkg/time/clock",
		"time",
		[]string{"example.org/alfa/pkg/time", "time"},
		"example.org/alfa/pkg/time",
		standard("time"),
	)
	assert.String(t, "time", keeper)
	assert.String(t, "stdlib", reason)
}

func TestKeeperIsTheImportSharingTheLongestPath(t *testing.T) {
	keeper, reason := imports.Keeper(
		"example.org/alfa/pkg/tool/goalfa/server",
		constant.ConstantPackage,
		[]string{
			"example.org/alfa/pkg/web/constant",
			"example.org/alfa/pkg/tool/goalfa/constant",
		},
		"example.org/alfa/pkg/web/constant",
		standard(),
	)
	assert.String(t, "example.org/alfa/pkg/tool/goalfa/constant", keeper)
	assert.String(t, "own", reason)
}

func TestKeeperOnATieIsTheDeeperPackage(t *testing.T) {
	keeper, reason := imports.Keeper(
		"example.org/alfa/pkg/tool/goalfa",
		constant.ConstantPackage,
		[]string{
			"example.org/alfa/pkg/constant",
			"example.org/alfa/pkg/monitor/constant",
		},
		"example.org/alfa/pkg/constant",
		standard(),
	)
	assert.String(t, "example.org/alfa/pkg/monitor/constant", keeper)
	assert.String(t, "deeper", reason)
}

func TestKeeperOnEqualDepthIsTheOneAlreadyPlain(t *testing.T) {
	keeper, reason := imports.Keeper(
		"example.org/alfa/pkg/tool/goalfa",
		constant.ConstantPackage,
		[]string{
			"example.org/alfa/pkg/web/constant",
			"example.org/alfa/pkg/generative/constant",
		},
		"example.org/alfa/pkg/web/constant",
		standard(),
	)
	assert.String(t, "example.org/alfa/pkg/web/constant", keeper)
	assert.String(t, "already plain", reason)
}

func TestKeeperOtherwiseIsTheFirstByPath(t *testing.T) {
	keeper, reason := imports.Keeper(
		"example.org/alfa/pkg/tool/goalfa",
		constant.ConstantPackage,
		[]string{
			"example.org/alfa/pkg/web/constant",
			"example.org/alfa/pkg/generative/constant",
		},
		"",
		standard(),
	)
	assert.String(t, "example.org/alfa/pkg/generative/constant", keeper)
	assert.String(t, "first by path", reason)
}

func TestKeeperOfVersionNamesIsNobody(t *testing.T) {
	keeper, reason := imports.Keeper(
		"example.org/alfa/pkg/cluster",
		"v1",
		[]string{"example.org/core/v1", "example.org/meta/v1"},
		"example.org/meta/v1",
		standard(),
	)
	assert.String(t, "", keeper)
	assert.String(t, "version name", reason)
}
