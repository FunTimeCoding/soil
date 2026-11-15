package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/test_event/constant"
	"github.com/funtimecoding/soil/pkg/test_event/unit/test_event_tester"
	"testing"
)

func TestRunTests(t *testing.T) {
	assert.Strings(
		t,
		[]string{"TestFails", "TestPanics", "TestPasses"},
		test_event_tester.Run(constant.FixtureMixedStream).Tests(),
	)
}

func TestRunPassed(t *testing.T) {
	assert.Strings(
		t,
		[]string{"TestPasses"},
		test_event_tester.Run(constant.FixtureMixedStream).Passed(),
	)
}

func TestRunFailedNamesPanicToo(t *testing.T) {
	assert.Strings(
		t,
		[]string{"TestFails", "TestPanics"},
		test_event_tester.Run(constant.FixtureMixedStream).Failed(),
	)
}

func TestRunPanicked(t *testing.T) {
	assert.Strings(
		t,
		[]string{"TestPanics"},
		test_event_tester.Run(constant.FixtureMixedStream).Panicked(),
	)
}

func TestRunBuildOutputCarriesCompilerLine(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			"# probe/broken [probe/broken.test]",
			"broken/broken.go:4:2: declared and not used: x",
		},
		test_event_tester.Run(constant.FixtureBuildFailureStream).BuildOutput(),
	)
}

func TestRunCompilerLineSkipsPackageHeader(t *testing.T) {
	assert.String(
		t,
		"broken/broken.go:4:2: declared and not used: x",
		test_event_tester.Run(constant.FixtureBuildFailureStream).CompilerLine(),
	)
}

func TestRunCompilerLineEmptyWithoutBuildOutput(t *testing.T) {
	assert.String(
		t,
		"",
		test_event_tester.Run(constant.FixtureMixedStream).CompilerLine(),
	)
}

func TestRunFailedPackages(t *testing.T) {
	assert.Strings(
		t,
		[]string{"probe/broken"},
		test_event_tester.Run(constant.FixtureBuildFailureStream).FailedPackages(),
	)
}

func TestRunFailedPackagesOncePerPackage(t *testing.T) {
	assert.Strings(
		t,
		[]string{"probe/good"},
		test_event_tester.Run(constant.FixtureMixedStream).FailedPackages(),
	)
}

func TestRunSetupPanicFailsPackageWithoutTest(t *testing.T) {
	r := test_event_tester.Run(constant.FixtureSetupPanicStream)
	assert.Strings(t, []string{"probe/setup"}, r.FailedPackages())
	assert.Count(t, 0, r.Failed())
	assert.Count(t, 0, r.Panicked())
}

func TestRunWithoutTests(t *testing.T) {
	r := test_event_tester.Run(constant.FixtureNoTestsStream)
	assert.Count(t, 0, r.Tests())
	assert.Count(t, 0, r.FailedPackages())
}
