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
