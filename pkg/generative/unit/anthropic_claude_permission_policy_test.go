package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"os"
	"path/filepath"
	"testing"
)

func TestPolicyAllowsALoopOfAllowlistedCommands(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(
			t,
			root,
			scratchpad,
			`for p in 7 8 9; do app run --rm alpha /alpha/render.py "/input/x.txt" "part:$p" /out/p$p.txt 2>&1 | head -1; done`,
		),
	)
}

func TestPolicyAllowsRemovingInsideScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(t, allows(t, root, scratchpad, `rm -f "tmp/work/x clean.md"`))
}

func TestPolicyAllowsAGlobInsideAScratchDirectory(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(t, allows(t, root, scratchpad, `rm -f tmp/work/*.md`))
}

func TestPolicyAllowsAVariableNamingTheScratchpad(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(
			t,
			root,
			scratchpad,
			fmt.Sprintf(`S=%s/work; mkdir -p "$S" && rm -rf "$S"`, scratchpad),
		),
	)
}

func TestPolicyAllowsRedirectingIntoScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(
			t,
			root,
			scratchpad,
			`grep -c x doc/guide/a.md > tmp/work/count.txt 2>/dev/null`,
		),
	)
}

func TestPolicyAllowsAnAllowlistedSubstitution(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(t, root, scratchpad, `echo "$(wc -l < doc/guide/a.md)"`),
	)
}

func TestPolicyAllowsCopyingIntoScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(t, allows(t, root, scratchpad, `cp doc/guide/a.md tmp/work/`))
}

func TestPolicyFollowsAChangeDirectory(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(t, allows(t, root, scratchpad, `cd tmp/work && rm -f x.md`))
}

func TestPolicyAllowsShellBuiltinsInsideALoop(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(
			t,
			root,
			scratchpad,
			"set -e\nls doc/guide | while IFS= read -r f; do [ \"$f\" = README.md ] && continue; printf '%s\\n' \"$f\"; done",
		),
	)
}

func TestPolicyRefusesAShellBuiltinTheAllowlistLacks(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `ls doc; shift`))
}

func TestPolicyAllowsATimeoutAroundAnAllowlistedCommand(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(
			t,
			root,
			scratchpad,
			`timeout -k 5 900 app run --rm alpha /alpha/count.py /input`,
		),
	)
}

func TestPolicyRefusesATimeoutAroundAnUnlistedCommand(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `timeout 900 python3 x.py`))
}

func TestPolicyRefusesAFieldSeparatorChangeThatOutlivesOneCommand(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `IFS=/; ls doc`))
}

func TestPolicyForgetsAVariableThatReadOverwrites(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(
		t,
		allows(
			t,
			root,
			scratchpad,
			`S=tmp/work/a; echo doc | read S; rm -rf "$S"`,
		),
	)
}

func TestPolicyRefusesRemovingOutsideScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `rm doc/guide/a.md`))
}

func TestPolicyRefusesAPathClimbingOutOfScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `rm tmp/../doc/guide/a.md`))
}

func TestPolicyRefusesTheScratchRootItself(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `rm -rf tmp`))
}

func TestPolicyRefusesAGlobOverTheScratchRoot(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `rm -rf tmp/*`))
}

func TestPolicyRefusesADotGlob(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `rm -rf tmp/work/.*`))
}

func TestPolicyRefusesASymlinkOutOfScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.FatalOnError(
		t,
		os.Symlink(
			filepath.Join(root, "doc"),
			filepath.Join(
				root,
				constant.FixtureScratchDirectory,
				"work",
				"link",
			),
		),
	)
	assert.False(t, allows(t, root, scratchpad, `rm -rf tmp/work/link/guide`))
}

func TestPolicyRefusesAnUnlistedCommand(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(
		t,
		allows(t, root, scratchpad, "python3 - <<'EOF'\nprint(1)\nEOF"),
	)
}

func TestPolicyRefusesAnUnlistedCommandInsideALoop(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(
		t,
		allows(t, root, scratchpad, `for f in a b; do python3 x.py "$f"; done`),
	)
}

func TestPolicyRefusesAnUnlistedSubstitution(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(
		t,
		allows(t, root, scratchpad, `echo "$(python3 -c 'print(1)')"`),
	)
}

func TestPolicyRefusesRedirectingOutsideScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `echo x > doc/guide/a.md`))
}

func TestPolicyRefusesAnUnknownVariable(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `rm -rf "$HOME/x"`))
}

func TestPolicyForgetsAVariableSetInsideALoop(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(
		t,
		allows(
			t,
			root,
			scratchpad,
			`S=tmp/work/a; for d in doc; do S=$d; done; rm -rf "$S"`,
		),
	)
}

func TestPolicyKeepsASubshellChangeDirectoryInside(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `(cd tmp/work) && rm -f x.md`))
}

func TestPolicyRefusesAFunctionShadowingAnAllowlistedName(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(
		t,
		allows(t, root, scratchpad, `grep() { rm -rf doc; }; grep x a.md`),
	)
}

func TestPolicyRefusesAPathOverride(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `PATH=/tmp/evil grep x a.md`))
}

func TestPolicyRefusesFindExec(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(
		t,
		allows(t, root, scratchpad, `find . -name x -exec rm {} \;`),
	)
}

func TestPolicyAllowsSedInPlaceLikeTheAllowlist(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(t, root, scratchpad, `gsed -i 's/a/b/' doc/guide/a.md`),
	)
}

func TestPolicyRefusesAnUnknownLeadingAssignment(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `FOO=bar app run --rm alpha x`))
}

func TestPolicyAllowsALocaleBeforeACommand(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(t, root, scratchpad, `LC_ALL=C grep -c x doc/guide/a.md`),
	)
}

func TestPolicyRefusesAFieldSeparatorBeforeAnythingButRead(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `IFS=/ ls doc`))
}

func TestPolicyRefusesAnUnquotedGlobOnAWriteCapableCommand(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `find doc -name *.md`))
}

func TestPolicyAllowsAQuotedGlobOnAWriteCapableCommand(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(t, allows(t, root, scratchpad, `find doc -name '*.md'`))
}

func TestPolicyRefusesGitAfterChangingDirectory(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `cd tmp/work && git diff`))
}

func TestPolicyAllowsGitAfterAChangeDirectoryThatStays(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(t, root, scratchpad, fmt.Sprintf(`cd %s && git diff`, root)),
	)
}

func TestPolicyUnwrapsABareXargs(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(t, root, scratchpad, `ls doc/guide | xargs grep -c x`),
	)
}

func TestPolicyRefusesAnXargsWithFlags(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(
		t,
		allows(t, root, scratchpad, `ls doc/guide | xargs -n1 grep -c x`),
	)
}

func TestPolicyRefusesACommandLookup(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `command -v python3`))
}

func TestPolicyAllowsAReadOnlyCommandOutsideTheAllowlist(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.True(
		t,
		allows(
			t,
			root,
			scratchpad,
			`diff doc/guide/a.md tmp/work/a.md; stat doc`,
		),
	)
}

func TestPolicyRefusesAnUnparseableCommand(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `for x in; do`))
}

func TestPolicyRefusesABackgroundJob(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `grep x a.md &`))
}

func TestPolicyRefusesMovingOutOfScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `mv doc/guide/a.md tmp/work/`))
}

func TestPolicyRefusesCopyingOutOfScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.False(t, allows(t, root, scratchpad, `cp tmp/work/a.md doc/guide/`))
}

func TestPolicyNamesNoBlockerForAnAllowedCommand(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.Strings(
		t,
		nil,
		blockers(t, root, scratchpad, `ls doc && rm -f tmp/work/x.md`),
	)
}

func TestPolicyNamesEveryBlockerNotOnlyTheFirst(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.Strings(
		t,
		[]string{"python3", "rm", "md5"},
		blockers(
			t,
			root,
			scratchpad,
			`python3 -c 'print(1)' && rm doc/guide/a.md; python3 -V | md5`,
		),
	)
}

func TestPolicyNamesTheBlockerInsideLoopsAndSubstitutions(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.Strings(
		t,
		[]string{"cmp"},
		blockers(t, root, scratchpad, `for f in a b; do cmp "$f" x; done`),
	)
	assert.Strings(
		t,
		[]string{"md5"},
		blockers(t, root, scratchpad, `echo "$(md5 doc/guide/a.md)"`),
	)
	assert.Strings(
		t,
		[]string{"shasum"},
		blockers(
			t,
			root,
			scratchpad,
			`if ls doc; then shasum doc/guide/a.md; fi`,
		),
	)
}

func TestPolicyWalksOnPastARefusal(t *testing.T) {
	root, scratchpad := newProject(t)
	both := []string{"cmp", "shasum"}
	assert.Strings(
		t,
		both,
		blockers(t, root, scratchpad, `if cmp a b; then shasum x; fi`),
	)
	assert.Strings(
		t,
		both,
		blockers(t, root, scratchpad, `while cmp a b; do shasum x; done`),
	)
	assert.Strings(
		t,
		both,
		blockers(t, root, scratchpad, `cmp a b || shasum x`),
	)
	assert.Strings(
		t,
		both,
		blockers(t, root, scratchpad, `echo "$(cmp a b)$(shasum x)"`),
	)
	assert.Strings(
		t,
		both,
		blockers(t, root, scratchpad, `for f in $(cmp a b); do shasum x; done`),
	)
	assert.Strings(
		t,
		both,
		blockers(t, root, scratchpad, `case $(cmp a b) in a) shasum x;; esac`),
	)
	assert.Strings(
		t,
		[]string{"(assignment before a command)", "cmp"},
		blockers(t, root, scratchpad, `X=1 cmp a b`),
	)
	assert.Strings(
		t,
		[]string{"(assignment to PATH)", "cmp"},
		blockers(t, root, scratchpad, `export PATH=/opt HOME=$(cmp a b)`),
	)
	assert.Strings(
		t,
		[]string{"(background job)", "cmp"},
		blockers(t, root, scratchpad, `cmp a b &`),
	)
	assert.Strings(
		t,
		[]string{"cmp", "(redirect outside scratch)"},
		blockers(t, root, scratchpad, `ls > "$(cmp a b)"`),
	)
	assert.Strings(
		t,
		[]string{"(redirect outside scratch)", "cmp"},
		blockers(t, root, scratchpad, `cmp a b > doc/list.txt`),
	)
}

func TestPolicyNamesTheSubcommandWhereRulesScopeOne(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.Strings(
		t,
		[]string{"git status"},
		blockers(t, root, scratchpad, `git status --short`),
	)
	assert.Strings(
		t,
		[]string{"app ps"},
		blockers(t, root, scratchpad, `app ps`),
	)
	assert.Strings(t, []string{"cmp"}, blockers(t, root, scratchpad, `cmp a b`))
}

func TestPolicyNamesGitAfterAChangeDirectory(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.Strings(
		t,
		[]string{"(git after cd)"},
		blockers(t, root, scratchpad, `cd doc && git diff`),
	)
}

func TestPolicyNamesAWriteOutsideScratch(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.Strings(
		t,
		[]string{"rm"},
		blockers(t, root, scratchpad, `rm doc/guide/a.md`),
	)
	assert.Strings(
		t,
		[]string{"cp"},
		blockers(t, root, scratchpad, `cp tmp/work/a.md doc/guide/`),
	)
	assert.Strings(
		t,
		[]string{"(redirect outside scratch)"},
		blockers(t, root, scratchpad, `ls doc > doc/list.txt`),
	)
}

func TestPolicyNamesStructuralBlocks(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.Strings(
		t,
		[]string{"(background job)"},
		blockers(t, root, scratchpad, `ls doc &`),
	)
	assert.Strings(
		t,
		[]string{"(assignment before a command)"},
		blockers(t, root, scratchpad, `X=1 ls doc`),
	)
	assert.Strings(
		t,
		[]string{"(assignment to PATH)"},
		blockers(t, root, scratchpad, `export PATH=/opt/bin; ls doc`),
	)
	assert.Strings(
		t,
		[]string{"(array assignment)"},
		blockers(t, root, scratchpad, `a=(1 2); ls doc`),
	)
	assert.Strings(
		t,
		[]string{"(command from an expansion)"},
		blockers(t, root, scratchpad, `$tool doc`),
	)
	assert.Strings(
		t,
		[]string{"(unsupported construct)"},
		blockers(t, root, scratchpad, `f() { ls doc; }`),
	)
	assert.Strings(
		t,
		[]string{"(unparsed)"},
		blockers(t, root, scratchpad, `ls doc (`),
	)
	assert.Strings(
		t,
		[]string{"(no command)"},
		blockers(t, root, scratchpad, `x=1`),
	)
	assert.Strings(
		t,
		[]string{"find"},
		blockers(t, root, scratchpad, `find doc -delete`),
	)
}

func TestPolicyNamesAWrapperItCannotStrip(t *testing.T) {
	root, scratchpad := newProject(t)
	assert.Strings(
		t,
		[]string{"xargs"},
		blockers(t, root, scratchpad, `ls doc | xargs -n1 cat`),
	)
	assert.Strings(
		t,
		[]string{"timeout"},
		blockers(t, root, scratchpad, `timeout 5`),
	)
	assert.Strings(
		t,
		[]string{"cmp"},
		blockers(t, root, scratchpad, `timeout 5 cmp a b`),
	)
}
