package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/service_tester"
	"path/filepath"
	"testing"
)

func TestRemoveParametersFollowsAChainThroughTheBatch(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("remove-parameters/src"),
	)
	r, e := testService().RemoveParameters(
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
	assertFormatted(t, mount)
	run := service_tester.ReadFixtureFile(t, d, "pkg/target/run.go")
	assert.StringContains(t, "func Run(\n\to *Option,\n) string {", run)
	assert.StringContains(t, "Mount(o.Name, 1)", run)
	caller := service_tester.ReadFixtureFile(t, d, "pkg/caller/main.go")
	assert.StringContains(t, "target.Run(o)", caller)
	assert.StringNotContains(t, "label", caller)
	assertFormatted(t, caller)
	assert.StringContains(
		t,
		"field Version is now only written",
		joinedConcerns(r),
	)
}

func TestRemoveParametersDropsEveryVariadicArgument(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("remove-parameters/src"),
	)
	r, e := testService().RemoveParameters(
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
		joinedConcerns(r),
	)
	assertFormatted(t, caller)
}

func TestRemoveParametersDryRunWritesNothing(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("remove-parameters/src"),
	)
	before := service_tester.ReadFixtureFile(t, d, "pkg/target/log.go")
	r, e := testService().RemoveParameters(
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

func TestRemoveParametersRefusesAParameterStillUsed(t *testing.T) {
	assertRemovalRefused(t, "parameter version of Used is still used", "Used")
}

func TestRemoveParametersRefusesPassingOnOutsideTheBatch(t *testing.T) {
	assertRemovalRefused(
		t,
		"parameter version of Forward is still used",
		"Forward",
	)
}

func TestRemoveParametersRefusesAFunctionUsedAsAValue(t *testing.T) {
	assertRemovalRefused(t, "Valued is used as a value", "Valued")
}

func TestRemoveParametersRefusesDroppingACall(t *testing.T) {
	assertRemovalRefused(
		t,
		"dropping the argument current() would remove a call",
		"Called",
	)
}

func TestRemoveParametersRefusesRemovingALocalThatCalls(t *testing.T) {
	assertRemovalRefused(
		t,
		"removing the unused local version would remove a call",
		"Local",
	)
}

func TestRemoveParametersRefusesAMethodSatisfyingAnInterface(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("remove-parameters-interface/src"),
	)
	r, e := testService().RemoveParameters(
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

func TestRemoveParametersRefusesAnUnknownParameter(t *testing.T) {
	d := testutil.PrepareTestPackage(
		t,
		service_tester.ServiceTestdata("remove-parameters/src"),
	)
	r, e := testService().RemoveParameters(
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

func TestRemoveParametersRewritesAReplacingModule(t *testing.T) {
	for _, full := range []bool{false, true} {
		library, user := prepareReplacedLibrary(t, "remove-parameters-replacer")
		s := replacingService(library, user)

		if full {
			s.UseFullLoad()
		}

		r, e := s.RemoveParameters(library, logValues(), false)
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

func TestRemoveParametersRefusesASideEffectInAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "remove-parameters-replacer")
	testutil.WriteFile(
		t,
		user,
		"pkg/effect/effect.go",
		"package effect\n\nimport \"other.test/lib\"\n\nfunc count() int {\n\treturn 1\n}\n\nfunc Use() string {\n\treturn lib.Log(\"a\", count())\n}\n",
	)
	source := service_tester.ReadFixtureFile(t, library, "lib.go")
	r, e := replacingService(library, user).RemoveParameters(
		library,
		logValues(),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		filepath.Join(user, "pkg/effect/effect.go"),
		joinedConcerns(r),
	)
	assert.String(
		t,
		source,
		service_tester.ReadFixtureFile(t, library, "lib.go"),
	)
}

func TestRemoveParametersRefusesAValueUseInAReplacingModule(t *testing.T) {
	library, user := prepareReplacedLibrary(t, "remove-parameters-replacer")
	testutil.WriteFile(
		t,
		user,
		"pkg/value/value.go",
		"package value\n\nimport \"other.test/lib\"\n\nvar Logger = lib.Log\n",
	)
	r, e := replacingService(library, user).RemoveParameters(
		library,
		logValues(),
		false,
	)
	assert.FatalOnError(t, e)
	testutil.AssertBlocked(t, r, 1)
	assert.StringContains(
		t,
		filepath.Join(user, "pkg/value/value.go"),
		joinedConcerns(r),
	)
}
