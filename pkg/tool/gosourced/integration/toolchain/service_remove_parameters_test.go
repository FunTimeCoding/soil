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
	o := service_tester.New(
		t,
		service_tester.WithFixture("remove-parameters/src"),
	)
	r, e := o.Symbol.RemoveParameters(
		t.Context(),
		o.Directory,
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
	mount := o.Read("pkg/target/mount.go")
	assert.StringContains(t, "\tname string,\n\tport int,\n) string {", mount)
	rewrite_tester.Formatted(t, mount)
	run := o.Read("pkg/target/run.go")
	assert.StringContains(t, "func Run(\n\to *Option,\n) string {", run)
	assert.StringContains(t, "Mount(o.Name, 1)", run)
	caller := o.Read("pkg/caller/main.go")
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
	o := service_tester.New(
		t,
		service_tester.WithFixture("remove-parameters/src"),
	)
	r, e := o.Symbol.RemoveParameters(
		t.Context(),
		o.Directory,
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
	caller := o.Read("pkg/caller/main.go")
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
	o := service_tester.New(
		t,
		service_tester.WithFixture("remove-parameters/src"),
	)
	before := o.Read("pkg/target/log.go")
	r, e := o.Symbol.RemoveParameters(
		t.Context(),
		o.Directory,
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
	assert.String(t, before, o.Read("pkg/target/log.go"))
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
	o := service_tester.New(
		t,
		service_tester.WithFixture("remove-parameters-interface/src"),
	)
	r, e := o.Symbol.RemoveParameters(
		t.Context(),
		o.Directory,
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
	o := service_tester.New(
		t,
		service_tester.WithFixture("remove-parameters/src"),
	)
	r, e := o.Symbol.RemoveParameters(
		t.Context(),
		o.Directory,
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
		library, o := module_tester.ReplacedLibrary(
			t,
			"remove-parameters-replacer",
		)

		if full {
			o.Workspace.UseFullLoad()
		}

		r, e := o.Symbol.RemoveParameters(
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
			system_tester.Read(t, library, "lib.go"),
		)
		assert.String(
			t,
			"package user\n\nimport \"other.test/lib\"\n\nfunc Use() string {\n\treturn lib.Log(\"a\")\n}\n",
			o.Read("pkg/user/user.go"),
		)
	}
}

func TestRemoveParametersRefusesSideEffectInReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "remove-parameters-replacer")
	system_tester.Write(
		t,
		o.Directory,
		"pkg/effect/effect.go",
		"package effect\n\nimport \"other.test/lib\"\n\nfunc count() int {\n\treturn 1\n}\n\nfunc Use() string {\n\treturn lib.Log(\"a\", count())\n}\n",
	)
	source := system_tester.Read(t, library, "lib.go")
	r, e := o.Symbol.RemoveParameters(
		t.Context(),
		library,
		module_tester.LogValues(),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		filepath.Join(o.Directory, "pkg/effect/effect.go"),
		rewrite_tester.JoinedConcerns(r),
	)
	assert.String(t, source, system_tester.Read(t, library, "lib.go"))
}

func TestRemoveParametersRefusesValueUseInReplacingModule(t *testing.T) {
	library, o := module_tester.ReplacedLibrary(t, "remove-parameters-replacer")
	system_tester.Write(
		t,
		o.Directory,
		"pkg/value/value.go",
		"package value\n\nimport \"other.test/lib\"\n\nvar Logger = lib.Log\n",
	)
	r, e := o.Symbol.RemoveParameters(
		t.Context(),
		library,
		module_tester.LogValues(),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		filepath.Join(o.Directory, "pkg/value/value.go"),
		rewrite_tester.JoinedConcerns(r),
	)
}
