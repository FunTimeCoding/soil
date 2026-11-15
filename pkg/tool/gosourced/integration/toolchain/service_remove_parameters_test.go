package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/module_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration/toolchain/rewrite_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/removal"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestRemoveParametersFollowsChainThroughBatch(t *testing.T) {
	d := service_tester.Prepare(t, "remove-parameters/src")
	r, e := service_tester.NewService(t).Symbol.RemoveParameters(
		t.Context(),
		d,
		[]*removal.Parameter{
			removal.NewParameter(
				"example/pkg/target",
				"Mount",
				"",
				[]string{"version"},
			),
			removal.NewParameter(
				"example/pkg/target",
				"Run",
				"",
				[]string{"version"},
			),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	mount := service_tester.ReadFixtureFile(t, d, "pkg/target/mount.go")
	assert.StringContains(t, "\tname string,\n\tport int,\n) string {", mount)
	rewrite_tester.Formatted(t, mount)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "func Run(\n\to *Option,\n) string {", run)
	assert.StringContains(t, "Mount(o.Name, 1)", run)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/main.go")
	assert.StringContains(t, "target.Run(o)", caller)
	assert.StringNotContains(t, "label", caller)
	rewrite_tester.Formatted(t, caller)
	assert.StringContains(
		t,
		"field Version is now only written",
		rewrite_tester.JoinedConcerns(r),
	)
}

func TestRemoveParametersDropsEveryVariadicArgument(t *testing.T) {
	d := service_tester.Prepare(t, "remove-parameters/src")
	r, e := service_tester.NewService(t).Symbol.RemoveParameters(
		t.Context(),
		d,
		[]*removal.Parameter{
			removal.NewParameter(
				"example/pkg/target",
				"Log",
				"",
				[]string{"values"},
			),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/main.go")
	assert.StringContains(t, `target.Log("a")`, caller)
	assert.StringContains(t, `target.Log("b")`, caller)
	assert.StringContains(
		t,
		"parameter values is no longer used",
		rewrite_tester.JoinedConcerns(r),
	)
	rewrite_tester.Formatted(t, caller)
}

func TestRemoveParametersDryRunWritesNothing(t *testing.T) {
	d := service_tester.Prepare(t, "remove-parameters/src")
	before := service_tester.ReadFixtureFile(t, d, "pkg/target/log.go")
	r, e := service_tester.NewService(t).Symbol.RemoveParameters(
		t.Context(),
		d,
		[]*removal.Parameter{
			removal.NewParameter(
				"example/pkg/target",
				"Log",
				"",
				[]string{"values"},
			),
		},
		true,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 0)
	assert.String(
		t,
		before,
		service_tester.ReadFixtureFile(t, d, "pkg/target/log.go"),
	)
}

func TestRemoveParametersRefusesParameterStillUsed(t *testing.T) {
	rewrite_tester.RemovalRefused(
		t,
		"parameter version of Used is still used",
		"Used",
	)
}

func TestRemoveParametersRefusesPassingOnOutsideBatch(t *testing.T) {
	rewrite_tester.RemovalRefused(
		t,
		"parameter version of Forward is still used",
		"Forward",
	)
}

func TestRemoveParametersRefusesFunctionUsedAsValue(t *testing.T) {
	rewrite_tester.RemovalRefused(t, "Valued is used as a value", "Valued")
}

func TestRemoveParametersRefusesDroppingCall(t *testing.T) {
	rewrite_tester.RemovalRefused(
		t,
		"dropping the argument current() would remove a call",
		"Called",
	)
}

func TestRemoveParametersRefusesRemovingLocalThatCalls(t *testing.T) {
	rewrite_tester.RemovalRefused(
		t,
		"removing the unused local version would remove a call",
		"Local",
	)
}

func TestRemoveParametersRefusesMethodSatisfyingInterface(t *testing.T) {
	d := service_tester.Prepare(t, "remove-parameters-interface/src")
	r, e := service_tester.NewService(t).Symbol.RemoveParameters(
		t.Context(),
		d,
		[]*removal.Parameter{
			removal.NewParameter(
				"example/pkg/target",
				"Save",
				"Store",
				[]string{"version"},
			),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(
		t,
		r,
		"Save satisfies example/pkg/target.Saver",
	)
}

func TestRemoveParametersRefusesUnknownParameter(t *testing.T) {
	d := service_tester.Prepare(t, "remove-parameters/src")
	r, e := service_tester.NewService(t).Symbol.RemoveParameters(
		t.Context(),
		d,
		[]*removal.Parameter{
			removal.NewParameter(
				"example/pkg/target",
				"Mount",
				"",
				[]string{"missing"},
			),
		},
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlockedContains(t, r, "has no parameter missing")
}

func TestRemoveParametersRewritesReplacingModule(t *testing.T) {
	for _, full := range []bool{false, true} {
		library, user := module_tester.ReplacedLibrary(
			t,
			"remove-parameters-replacer",
		)
		s := module_tester.ReplacingService(t, library, user)

		if full {
			s.Workspace.UseFullLoad()
		}

		r, e := s.Symbol.RemoveParameters(
			t.Context(),
			library,
			module_tester.LogValues(),
			false,
		)
		assert.FatalOnError(t, e)
		testutil.AssertBlocked(t, r, 0)
		assert.String(
			t,
			"package lib\n\nfunc Log(\n\tprefix string,\n) string {\n\treturn prefix\n}\n",
			service_tester.ReadFixtureFile(t, library, "lib.go"),
		)
		assert.String(
			t,
			"package user\n\nimport \"other.test/lib\"\n\nfunc Use() string {\n\treturn lib.Log(\"a\")\n}\n",
			service_tester.ReadFixtureFile(t, user, "pkg/user/user.go"),
		)
	}
}

func TestRemoveParametersRefusesSideEffectInReplacingModule(t *testing.T) {
	library, user := module_tester.ReplacedLibrary(
		t,
		"remove-parameters-replacer",
	)
	system_tester.Write(
		t,
		user,
		"pkg/effect/effect.go",
		"package effect\n\nimport \"other.test/lib\"\n\nfunc count() int {\n\treturn 1\n}\n\nfunc Use() string {\n\treturn lib.Log(\"a\", count())\n}\n",
	)
	source := service_tester.ReadFixtureFile(t, library, "lib.go")
	r, e := module_tester.ReplacingService(t, library, user).Symbol.RemoveParameters(
		t.Context(),
		library,
		module_tester.LogValues(),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		filepath.Join(user, "pkg/effect/effect.go"),
		rewrite_tester.JoinedConcerns(r),
	)
	assert.String(
		t,
		source,
		service_tester.ReadFixtureFile(t, library, "lib.go"),
	)
}

func TestRemoveParametersRefusesValueUseInReplacingModule(t *testing.T) {
	library, user := module_tester.ReplacedLibrary(
		t,
		"remove-parameters-replacer",
	)
	system_tester.Write(
		t,
		user,
		"pkg/value/value.go",
		"package value\n\nimport \"other.test/lib\"\n\nvar Logger = lib.Log\n",
	)
	r, e := module_tester.ReplacingService(t, library, user).Symbol.RemoveParameters(
		t.Context(),
		library,
		module_tester.LogValues(),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		filepath.Join(user, "pkg/value/value.go"),
		rewrite_tester.JoinedConcerns(r),
	)
}
