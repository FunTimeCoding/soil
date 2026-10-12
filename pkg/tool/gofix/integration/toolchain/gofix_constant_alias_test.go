package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gofix"
	"github.com/funtimecoding/soil/pkg/tool/gofix/integration/module_tester"
	"path/filepath"
	"testing"
)

func TestConstantAliasSuggestion(t *testing.T) {
	directory := module_tester.ConstantAlias(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunImportAliasFixWithDirectory([]string{"./..."}, directory, false, r)
	t.Run(
		"OwnIsPlainForeignCarriesItsDomain",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/alfa/own.go",
				"suggest de-alias own → constant (testmodule/pkg/alfa/constant)",
			)
			assertPlanned(
				t,
				r.Entries,
				"pkg/alfa/own.go",
				"suggest alias constant → bravoConstant (testmodule/pkg/bravo/constant)",
			)
		},
	)
	t.Run(
		"ForeignCarriesItsDomainEvenAlone",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/charlie/lone.go",
				"suggest alias constant → bravoConstant (testmodule/pkg/bravo/constant)",
			)
		},
	)
	t.Run(
		"LayoutRootCarriesTheModuleName",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/delta/root.go",
				"suggest alias constant → testmodule (testmodule/pkg/constant)",
			)
		},
	)
	t.Run(
		"DeeperOwnIsPlain",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/system/linux/debian/deeper.go",
				"suggest de-alias linux → constant (testmodule/pkg/system/linux/constant)",
			)
			assertPlanned(
				t,
				r.Entries,
				"pkg/system/linux/debian/deeper.go",
				"suggest alias constant → systemConstant (testmodule/pkg/system/constant)",
			)
		},
	)
	t.Run(
		"RefusesALocalName",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/juliet/local.go",
				"cannot de-alias own → constant (local collision with \"constant\")",
			)
		},
	)
	t.Run(
		"RefusesANameAnotherImportHolds",
		func(t *testing.T) {
			assertPlanned(
				t,
				r.Entries,
				"pkg/kilo/held.go",
				"cannot alias constant → bravoConstant (import collision with testmodule/pkg/kilo/vocabulary)",
			)
		},
	)
	t.Run(
		"ConformingFileIsQuiet",
		func(t *testing.T) {
			for _, c := range r.Entries {
				if c.Path == "pkg/golf/conforming.go" {
					t.Errorf("unexpected concern %q", c.Text)
				}
			}
		},
	)
	t.Run(
		"SuggestionsNeverEdit",
		func(t *testing.T) {
			assert.String(
				t,
				"package charlie\n\nimport \"testmodule/pkg/bravo/constant\"\n\nfunc Lone() string {\n\treturn constant.Bravo\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/charlie/lone.go"),
				),
			)
		},
	)
}

func TestConstantRealias(t *testing.T) {
	directory := module_tester.ConstantAlias(t)
	r := output.NewResultsWithDirectory(directory)
	gofix.RunRealiasFixWithDirectory([]string{"./..."}, directory, false, r)
	t.Run(
		"OwnIsPlainForeignCarriesItsDomain",
		func(t *testing.T) {
			assert.String(
				t,
				"package alfa\n\nimport (\n\t\"testmodule/pkg/alfa/constant\"\n\tbravoConstant \"testmodule/pkg/bravo/constant\"\n)\n\nfunc Own() string {\n\treturn constant.Alfa + bravoConstant.Bravo\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/alfa/own.go"),
				),
			)
		},
	)
	t.Run(
		"ForeignCarriesItsDomainEvenAlone",
		func(t *testing.T) {
			assert.String(
				t,
				"package charlie\n\nimport bravoConstant \"testmodule/pkg/bravo/constant\"\n\nfunc Lone() string {\n\treturn bravoConstant.Bravo\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/charlie/lone.go"),
				),
			)
		},
	)
	t.Run(
		"LayoutRootCarriesTheModuleName",
		func(t *testing.T) {
			assert.String(
				t,
				"package delta\n\nimport testmodule \"testmodule/pkg/constant\"\n\nfunc Root() string {\n\treturn testmodule.Root\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/delta/root.go"),
				),
			)
		},
	)
	t.Run(
		"DeeperOwnIsPlain",
		func(t *testing.T) {
			assert.String(
				t,
				"package debian\n\nimport (\n\tsystemConstant \"testmodule/pkg/system/constant\"\n\t\"testmodule/pkg/system/linux/constant\"\n)\n\nfunc Deeper() string {\n\treturn systemConstant.System + constant.Linux\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(
						directory,
						"pkg/system/linux/debian/deeper.go",
					),
				),
			)
		},
	)
	t.Run(
		"RefusedOwnKeepsItsAliasForeignStillMoves",
		func(t *testing.T) {
			assert.String(
				t,
				"package juliet\n\nimport (\n\tbravoConstant \"testmodule/pkg/bravo/constant\"\n\town \"testmodule/pkg/juliet/constant\"\n)\n\nfunc Local() string {\n\tconstant := own.Juliet\n\n\treturn constant + bravoConstant.Bravo\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/juliet/local.go"),
				),
			)
		},
	)
	t.Run(
		"ConformingFileIsUntouched",
		func(t *testing.T) {
			assert.String(
				t,
				"package golf\n\nimport (\n\tbravoConstant \"testmodule/pkg/bravo/constant\"\n\t\"testmodule/pkg/golf/constant\"\n)\n\nfunc Conforming() string {\n\treturn bravoConstant.Bravo + constant.Golf\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/golf/conforming.go"),
				),
			)
		},
	)
	t.Run(
		"ModuleStillBuilds",
		func(t *testing.T) {
			assertBuilds(t, directory)
		},
	)
}
