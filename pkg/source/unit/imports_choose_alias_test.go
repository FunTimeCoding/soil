package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"testing"
)

func TestChooseAliasPrefersTheBareSubsystem(t *testing.T) {
	assert.String(
		t,
		"web",
		imports.ChooseAlias(
			"example.org/alfa/pkg/web/constant",
			"example.org/alfa",
			constant.ConstantPackage,
			taken(),
		),
	)
}

func TestChooseAliasFallsBackToSubsystemAndName(t *testing.T) {
	assert.String(
		t,
		"stringsConstant",
		imports.ChooseAlias(
			"example.org/alfa/pkg/strings/constant",
			"example.org/alfa",
			constant.ConstantPackage,
			taken("strings"),
		),
	)
}

func TestChooseAliasFallbackCamelsASnakeName(t *testing.T) {
	assert.String(
		t,
		"clientMockClient",
		imports.ChooseAlias(
			"example.org/alfa/pkg/tool/goalfa/client/mock_client",
			"example.org/alfa",
			constant.MockPackage,
			taken("client"),
		),
	)
}

func TestChooseAliasAtTheLayoutRootIsModulePlusName(t *testing.T) {
	assert.String(
		t,
		"alfaErrors",
		imports.ChooseAlias(
			"example.org/alfa/pkg/errors",
			"example.org/alfa",
			"errors",
			taken("errors"),
		),
	)
}

func TestChooseAliasAtTheLayoutRootLeavesTheBareModuleToConstant(t *testing.T) {
	assert.String(
		t,
		"alfa",
		imports.ChooseAlias(
			"example.org/alfa/pkg/constant",
			"example.org/alfa",
			constant.ConstantPackage,
			taken(constant.ConstantPackage),
		),
	)
}

func TestChooseAliasGivesUpWhenBothAreTaken(t *testing.T) {
	assert.String(
		t,
		"",
		imports.ChooseAlias(
			"example.org/alfa/pkg/strings/constant",
			"example.org/alfa",
			constant.ConstantPackage,
			taken("strings", "stringsConstant"),
		),
	)
}
