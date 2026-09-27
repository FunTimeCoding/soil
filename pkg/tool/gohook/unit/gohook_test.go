package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git/changed"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	gohookConstant "github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"github.com/funtimecoding/soil/pkg/tool/gohook/job"
	"github.com/funtimecoding/soil/pkg/tool/gohook/option"
	"path/filepath"
	"testing"
)

func TestLoadToolPath(t *testing.T) {
	root := t.TempDir()
	write(
		root,
		gohookConstant.ToolFile,
		"pre-push:\n  - paths: ['**/*.go']\n    run: task lint\n  - run: task always\n",
	)
	c := configuration.Load(root)
	assert.String(t, filepath.Join(root, "strata/tool/gohook.yaml"), c.Path)
	assert.Strings(t, []string{"pre-push"}, c.HookNames())
	assert.Count(t, 2, c.Hooks["pre-push"])
	assert.String(t, "task lint", c.Hooks["pre-push"][0].Run)
	assert.Strings(t, []string{"**/*.go"}, c.Hooks["pre-push"][0].Paths)
	assert.Count(t, 0, c.Hooks["pre-push"][1].Paths)
}

func TestLoadRootBeatsTool(t *testing.T) {
	root := t.TempDir()
	write(root, gohookConstant.RootFile, "pre-push:\n  - run: root\n")
	write(root, gohookConstant.ToolFile, "pre-push:\n  - run: tool\n")
	c := configuration.Load(root)
	assert.String(t, filepath.Join(root, ".gohook.yaml"), c.Path)
	assert.String(t, "root", c.Hooks["pre-push"][0].Run)
}

func TestLoadMissingPanics(t *testing.T) {
	defer func() { assert.NotNil(t, recover()) }()
	configuration.Load(t.TempDir())
}

func TestLoadUnknownHookPanics(t *testing.T) {
	root := t.TempDir()
	write(root, gohookConstant.RootFile, "pre-psuh:\n  - run: x\n")
	defer func() { assert.NotNil(t, recover()) }()
	configuration.Load(root)
}

func TestLoadMissingRunPanics(t *testing.T) {
	root := t.TempDir()
	write(root, gohookConstant.RootFile, "pre-push:\n  - paths: [a/]\n")
	defer func() { assert.NotNil(t, recover()) }()
	configuration.Load(root)
}

func TestHookNamesSorted(t *testing.T) {
	root := t.TempDir()
	write(
		root,
		gohookConstant.RootFile,
		"pre-push:\n  - run: a\npre-commit:\n  - run: b\n",
	)
	assert.Strings(
		t,
		[]string{"pre-commit", "pre-push"},
		configuration.Load(root).HookNames(),
	)
}

func TestCommandSubstitutesArguments(t *testing.T) {
	j := job.New("gocommit check {1}")
	assert.String(
		t,
		"gocommit check .git/COMMIT_EDITMSG",
		j.Command([]string{".git/COMMIT_EDITMSG"}),
	)
}

func TestCommandSubstitutesAll(t *testing.T) {
	j := job.New("echo {@}")
	assert.String(t, "echo a b", j.Command([]string{"a", "b"}))
	assert.String(t, "echo ", j.Command(nil))
}

func TestCommandWithoutPlaceholders(t *testing.T) {
	assert.String(t, "task lint", job.New("task lint").Command([]string{"x"}))
}

func TestCommandMissingArgumentStaysLiteral(t *testing.T) {
	assert.String(t, "echo {2}", job.New("echo {2}").Command([]string{"a"}))
}

func TestMatchesPrefix(t *testing.T) {
	j := job.New("task lint-manifest", "strata/manifest/")
	assert.True(t, j.Matches([]string{"strata/manifest/deploy.yaml"}))
	assert.False(t, j.Matches([]string{"strata/manifests/deploy.yaml"}))
	assert.False(t, j.Matches([]string{"pkg/strata/manifest/x.yaml"}))
}

func TestMatchesGlob(t *testing.T) {
	j := job.New("task lint", "**/*.go")
	assert.True(t, j.Matches([]string{"pkg/alfa/bravo.go"}))
	assert.True(t, j.Matches([]string{"main.go"}))
	assert.False(t, j.Matches([]string{"pkg/alfa/bravo.yaml"}))
	assert.True(t, job.New("x", "*.md").Matches([]string{"README.md"}))
	assert.False(t, job.New("x", "*.md").Matches([]string{"doc/README.md"}))
}

func TestMatchesExact(t *testing.T) {
	j := job.New("go mod tidy", "go.mod")
	assert.True(t, j.Matches([]string{"go.mod"}))
	assert.False(t, j.Matches([]string{"pkg/go.mod"}))
}

func TestMatchesAnyPattern(t *testing.T) {
	j := job.New("x", "go.mod", "**/*.go")
	assert.True(t, j.Matches([]string{"pkg/a.go"}))
	assert.True(t, j.Matches([]string{"go.mod"}))
	assert.False(t, j.Matches([]string{"README.md"}))
}

func TestMatchesWithoutPaths(t *testing.T) {
	j := job.New("task always")
	assert.True(t, j.Matches(nil))
	assert.True(t, j.Matches([]string{"anything"}))
}

func TestMatchesNoChanges(t *testing.T) {
	assert.False(t, job.New("x", "**/*.go").Matches(nil))
}

func TestNewEntries(t *testing.T) {
	assert.Strings(
		t,
		[]string{"c.go"},
		gohook.NewEntries(
			[]string{"a.go", "b.go"},
			[]string{"a.go", "b.go", "c.go"},
		),
	)
	assert.Count(t, 0, gohook.NewEntries([]string{"a.go"}, []string{"a.go"}))
	assert.Count(t, 0, gohook.NewEntries([]string{"a.go"}, nil))
	assert.Strings(
		t,
		[]string{"a.go"},
		gohook.NewEntries(nil, []string{"a.go"}),
	)
}

func TestPushRefsUpdate(t *testing.T) {
	r := changed.NewPushRefs(
		"",
		"refs/heads/main abc123 refs/heads/main def456\n",
	)
	assert.False(t, r.None)
	assert.False(t, r.All)
	assert.String(t, "def456..abc123", r.String())
}

func TestPushRefsDeletionOnly(t *testing.T) {
	r := changed.NewPushRefs(
		"",
		join.Space(
			"(delete)",
			gitConstant.ZeroHash,
			"refs/heads/old",
			"abc123\n",
		),
	)
	assert.True(t, r.None)
	assert.String(t, "nothing pushed", r.String())
	assert.Count(t, 0, r.Files(""))
}

func TestPushRefsNewBranchFallsBackToUpstream(t *testing.T) {
	r := changed.NewPushRefs(
		t.TempDir(),
		join.Space(
			"refs/heads/feature",
			"abc123",
			"refs/heads/feature",
			gitConstant.ZeroHash,
		),
	)
	assert.False(t, r.None)
	assert.True(t, r.All)
}

func TestPushRefsSkipsDeletionBeforeUpdate(t *testing.T) {
	r := changed.NewPushRefs(
		"",
		join.Empty(
			join.Space(
				"(delete)",
				gitConstant.ZeroHash,
				"refs/heads/old",
				"abc123\n",
			),
			"refs/heads/main 111 refs/heads/main 222\n",
		),
	)
	assert.String(t, "222..111", r.String())
}

func TestPushRefsEmptyInputUsesUpstream(t *testing.T) {
	assert.True(t, changed.NewPushRefs(t.TempDir(), "").All)
}

func TestResolveRangeExplicitWinsOverInput(t *testing.T) {
	o := option.New()
	o.Hook = "pre-push"
	o.Base = "a"
	o.Head = "b"
	o.Input = "refs/heads/main x refs/heads/main y\n"
	assert.String(t, "a..b", gohook.ResolveRange("", o).String())
}

func TestResolveRangeInputOnlyForPrePush(t *testing.T) {
	o := option.New()
	o.Hook = "post-merge"
	o.Input = "refs/heads/main x refs/heads/main y\n"
	assert.True(t, gohook.ResolveRange("", o).All)
}
