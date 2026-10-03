package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/installed"
	"testing"
)

func TestLinkReadsTheStampFromLinkerFlags(t *testing.T) {
	b := installed.New("tool", "/bin/tool", "")
	b.Link(
		`"-X main.Version=v1.2.3 -X main.GitHash=abc1234 -X main.BuildDate=2026-01-01T00:00:00Z -X github.com/funtimecoding/soil/pkg/stamp/constant.Module=example.com/tool -X github.com/funtimecoding/soil/pkg/stamp/constant.Dirty=1"`,
	)
	assert.Any(
		t,
		&installed.Binary{
			Name:    "tool",
			Path:    "/bin/tool",
			Version: "v1.2.3",
			Hash:    "abc1234",
			Module:  "example.com/tool",
			Dirty:   true,
			Modules: map[string]string{},
		},
		b,
	)
}

func TestACleanStampIsNotDirty(t *testing.T) {
	b := installed.New("tool", "/bin/tool", "")
	b.Link("-X github.com/funtimecoding/soil/pkg/stamp/constant.Dirty=0")
	assert.Boolean(t, false, b.Dirty)
}

func TestTheMainPackageOfAFilePathBuildFollowsTheCommandDirectory(t *testing.T) {
	b := installed.New("golint", "/bin/golint", constant.CommandLineArguments)
	assert.String(
		t,
		"example.com/tool/cmd/golint",
		b.MainPackage("example.com/tool"),
	)
}

func TestTheMainPackageOfAnInstalledModuleIsItsOwnPath(t *testing.T) {
	b := installed.New("golint", "/bin/golint", "example.com/tool/cmd/golint")
	assert.String(
		t,
		"example.com/tool/cmd/golint",
		b.MainPackage("other.com/x"),
	)
}

func TestAChangedSourceFileTouchesItsPackage(t *testing.T) {
	assert.Boolean(
		t,
		true,
		installed.Touches(
			[]string{"pkg/argument/parse.go"},
			map[string]bool{"pkg/argument": true},
		),
	)
}

func TestAChangedTestFileDoesNotTouchTheBinary(t *testing.T) {
	assert.Boolean(
		t,
		false,
		installed.Touches(
			[]string{"cmd/golint/main_test.go"},
			map[string]bool{"cmd/golint": true},
		),
	)
}

func TestAChangeOutsideTheDependenciesDoesNotTouch(t *testing.T) {
	assert.Boolean(
		t,
		false,
		installed.Touches(
			[]string{"pkg/other/file.go"},
			map[string]bool{"pkg/argument": true},
		),
	)
}

func TestAModuleAtANewVersionHasMoved(t *testing.T) {
	assert.Boolean(
		t,
		true,
		installed.Moved(
			map[string]string{"example.com/a": "v1.0.0"},
			map[string]string{"example.com/a": "v1.1.0"},
		),
	)
}

func TestAModuleTheRepositoryNoLongerUsesHasNotMoved(t *testing.T) {
	assert.Boolean(
		t,
		false,
		installed.Moved(
			map[string]string{"example.com/gone": "v1.0.0"},
			map[string]string{"example.com/a": "v1.1.0"},
		),
	)
}
