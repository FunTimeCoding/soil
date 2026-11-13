package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/test_event/constant"
	"testing"
)

func TestRunTests(t *testing.T) {
	assert.Strings(
		t,
		[]string{"TestFails", "TestPanics", "TestPasses"},
		run(constant.FixtureMixedStream).Tests(),
	)
}

func TestRunPassed(t *testing.T) {
	assert.Strings(
		t,
		[]string{"TestPasses"},
		run(constant.FixtureMixedStream).Passed(),
	)
}

func TestRunFailedNamesPanicToo(t *testing.T) {
	assert.Strings(
		t,
		[]string{"TestFails", "TestPanics"},
		run(constant.FixtureMixedStream).Failed(),
	)
}

func TestRunPanicked(t *testing.T) {
	assert.Strings(
		t,
		[]string{"TestPanics"},
		run(constant.FixtureMixedStream).Panicked(),
	)
}

func TestRunBuildOutputCarriesCompilerLine(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			"# probe/broken [probe/broken.test]",
			"broken/broken.go:4:2: declared and not used: x",
		},
		run(constant.FixtureBuildFailureStream).BuildOutput(),
	)
}

func TestRunFailedPackages(t *testing.T) {
	assert.Strings(
		t,
		[]string{"probe/broken"},
		run(constant.FixtureBuildFailureStream).FailedPackages(),
	)
}

func TestRunFailedPackagesOncePerPackage(t *testing.T) {
	assert.Strings(
		t,
		[]string{"probe/good"},
		run(constant.FixtureMixedStream).FailedPackages(),
	)
}

func TestRunSetupPanicFailsPackageWithoutTest(t *testing.T) {
	r := run(constant.FixtureSetupPanicStream)
	assert.Strings(t, []string{"probe/setup"}, r.FailedPackages())
	assert.Count(t, 0, r.Failed())
	assert.Count(t, 0, r.Panicked())
}

func TestRunWithoutTests(t *testing.T) {
	r := run(constant.FixtureNoTestsStream)
	assert.Count(t, 0, r.Tests())
	assert.Count(t, 0, r.FailedPackages())
}
