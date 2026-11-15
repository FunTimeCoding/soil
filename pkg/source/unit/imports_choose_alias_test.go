package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/source/unit/import_tester"
	"testing"
)

func TestChooseAliasPrefersBareSubsystem(t *testing.T) {
	assert.String(
		t,
		"web",
		imports.ChooseAlias(
			"example.org/alfa/pkg/web/constant",
			"example.org/alfa",
			constant.ConstantPackage,
			import_tester.Among(),
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
			import_tester.Among("strings"),
		),
	)
}

func TestChooseAliasFallbackCamelsSnakeName(t *testing.T) {
	assert.String(
		t,
		"clientMockClient",
		imports.ChooseAlias(
			"example.org/alfa/pkg/tool/goalfa/client/mock_client",
			"example.org/alfa",
			constant.MockPackage,
			import_tester.Among("client"),
		),
	)
}

func TestChooseAliasKeepsBareSnakeSubsystem(t *testing.T) {
	assert.String(
		t,
		"bravo_charlie",
		imports.ChooseAlias(
			"example.org/alfa/pkg/bravo_charlie/constant",
			"example.org/alfa",
			constant.ConstantPackage,
			import_tester.Among(),
		),
	)
}

func TestChooseAliasFallbackCamelsSnakeSubsystem(t *testing.T) {
	assert.String(
		t,
		"bravoCharlieConstant",
		imports.ChooseAlias(
			"example.org/alfa/pkg/bravo_charlie/constant",
			"example.org/alfa",
			constant.ConstantPackage,
			import_tester.Among("bravo_charlie"),
		),
	)
}

func TestChooseAliasAtLayoutRootIsModulePlusName(t *testing.T) {
	assert.String(
		t,
		"alfaErrors",
		imports.ChooseAlias(
			"example.org/alfa/pkg/errors",
			"example.org/alfa",
			"errors",
			import_tester.Among("errors"),
		),
	)
}

func TestChooseAliasAtLayoutRootLeavesBareModuleToConstant(t *testing.T) {
	assert.String(
		t,
		"alfa",
		imports.ChooseAlias(
			"example.org/alfa/pkg/constant",
			"example.org/alfa",
			constant.ConstantPackage,
			import_tester.Among(constant.ConstantPackage),
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
			import_tester.Among("strings", "stringsConstant"),
		),
	)
}
