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

func TestDefaultRewritesDecidedCollisionKinds(t *testing.T) {
	directory := module_tester.CollisionKinds(t)
	r := output.NewResultsWithDirectory(directory)
	module_tester.RunDefault(
		t,
		module_tester.Options(t, directory, false, false, "./..."),
		r,
	)
	t.Run(
		"RootPackageTakesModulePlusName",
		func(t *testing.T) {
			assert.String(
				t,
				"package echo\n\nimport (\n\t\"errors\"\n\ttestmoduleErrors \"testmodule/pkg/errors\"\n)\n\nfunc Fail() error {\n\treturn errors.New(testmoduleErrors.Echo())\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/echo/fail.go"),
				),
			)
			assertResult(
				t,
				filterApplied(r.Entries),
				"pkg/echo/fail.go",
				"re-aliased library → testmoduleErrors (testmodule/pkg/errors)",
			)
		},
	)
	t.Run(
		"RootPackageBesideAKeeperSwapWaits",
		func(t *testing.T) {
			assert.String(
				t,
				"package service\n\nimport (\n\t\"testmodule/pkg/face\"\n\tgolfFace \"testmodule/pkg/golf/face\"\n)\n\nfunc Service(face.Reporter, golfFace.Client) {}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/golf/service/service.go"),
				),
			)
		},
	)
	t.Run(
		"VersionPairTakesItsGroup",
		func(t *testing.T) {
			assert.String(
				t,
				"package alfa\n\nimport (\n\tcore \"testmodule/pkg/core/v1\"\n\tmeta \"testmodule/pkg/meta/v1\"\n)\n\nfunc Pod() (*core.Pod, *meta.Object) {\n\treturn &core.Pod{}, &meta.Object{}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/alfa/pod.go"),
				),
			)
			assertResult(
				t,
				filterApplied(r.Entries),
				"pkg/alfa/pod.go",
				"re-aliased v1 → core (testmodule/pkg/core/v1)",
			)
		},
	)
	t.Run(
		"NumberedFallbackSwapsBack",
		func(t *testing.T) {
			assert.String(
				t,
				"package lower\n\nimport (\n\t\"testmodule/pkg/bravo/chain\"\n\tcharlie \"testmodule/pkg/charlie/chain\"\n)\n\nfunc Lower(*charlie.Link) *chain.Chain {\n\treturn &chain.Chain{}\n}\n",
				testutil.ReadFile(
					t,
					filepath.Join(directory, "pkg/bravo/lower/lower.go"),
				),
			)
			assertResult(
				t,
				filterApplied(r.Entries),
				"pkg/bravo/lower/lower.go",
				"de-aliased chain1 → chain (testmodule/pkg/bravo/chain)",
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

func TestLoserWithoutNameHoldsItsGroup(t *testing.T) {
	directory := module_tester.CollisionKinds(t)
	gofix.RunConstantAliasFix(
		[]string{"./..."},
		directory,
		false,
		output.NewResultsWithDirectory(directory),
	)
	assert.String(
		t,
		"package held\n\nimport (\n\t\"testmodule\"\n\tchain1 \"testmodule/pkg/delta/chain\"\n)\n\nfunc Held(*chain.Root) *chain1.Delta {\n\treturn &chain1.Delta{}\n}\n",
		testutil.ReadFile(t, filepath.Join(directory, "pkg/delta/held/held.go")),
	)
}
