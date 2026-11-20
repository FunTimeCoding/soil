package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/unit/policy_tester"
	"os"
	"path/filepath"
	"testing"
)

func TestPolicyAllowsLoopOfAllowlistedCommands(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(
		t,
		p.Allows(
			`for p in 7 8 9; do app run --rm alpha /alpha/render.py "/input/x.txt" "part:$p" /out/p$p.txt 2>&1 | head -1; done`,
		),
	)
}

func TestPolicyAllowsRemovingInsideScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`rm -f "tmp/work/x clean.md"`))
}

func TestPolicyAllowsGlobInsideScratchDirectory(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`rm -f tmp/work/*.md`))
}

func TestPolicyAllowsVariableNamingScratchpad(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(
		t,
		p.Allows(
			fmt.Sprintf(`S=%s/work; mkdir -p "$S" && rm -rf "$S"`, p.Scratchpad),
		),
	)
}

func TestPolicyAllowsRedirectingIntoScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(
		t,
		p.Allows(`grep -c x doc/guide/a.md > tmp/work/count.txt 2>/dev/null`),
	)
}

func TestPolicyAllowsAllowlistedSubstitution(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`echo "$(wc -l < doc/guide/a.md)"`))
}

func TestPolicyAllowsCopyingIntoScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`cp doc/guide/a.md tmp/work/`))
}

func TestPolicyFollowsChangeDirectory(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`cd tmp/work && rm -f x.md`))
}

func TestPolicyAllowsShellBuiltinsInsideLoop(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(
		t,
		p.Allows(
			"set -e\nls doc/guide | while IFS= read -r f; do [ \"$f\" = README.md ] && continue; printf '%s\\n' \"$f\"; done",
		),
	)
}

func TestPolicyRefusesShellBuiltinAllowlistLacks(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`ls doc; shift`))
}

func TestPolicyAllowsTimeoutAroundAllowlistedCommand(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(
		t,
		p.Allows(`timeout -k 5 900 app run --rm alpha /alpha/count.py /input`),
	)
}

func TestPolicyRefusesTimeoutAroundUnlistedCommand(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`timeout 900 python3 x.py`))
}

func TestPolicyRefusesFieldSeparatorChangeThatOutlivesOneCommand(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`IFS=/; ls doc`))
}

func TestPolicyForgetsVariableThatReadOverwrites(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`S=tmp/work/a; echo doc | read S; rm -rf "$S"`))
}

func TestPolicyRefusesRemovingOutsideScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`rm doc/guide/a.md`))
}

func TestPolicyRefusesPathClimbingOutOfScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`rm tmp/../doc/guide/a.md`))
}

func TestPolicyRefusesScratchRootItself(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`rm -rf tmp`))
}

func TestPolicyRefusesGlobOverScratchRoot(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`rm -rf tmp/*`))
}

func TestPolicyRefusesDotGlob(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`rm -rf tmp/work/.*`))
}

func TestPolicyRefusesSymlinkOutOfScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.FatalOnError(
		t,
		os.Symlink(
			filepath.Join(p.Root, "doc"),
			filepath.Join(
				p.Root,
				constant.FixtureScratchDirectory,
				"work",
				"link",
			),
		),
	)
	assert.False(t, p.Allows(`rm -rf tmp/work/link/guide`))
}

func TestPolicyRefusesUnlistedCommand(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows("python3 - <<'EOF'\nprint(1)\nEOF"))
}

func TestPolicyRefusesUnlistedCommandInsideLoop(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`for f in a b; do python3 x.py "$f"; done`))
}

func TestPolicyRefusesUnlistedSubstitution(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`echo "$(python3 -c 'print(1)')"`))
}

func TestPolicyRefusesRedirectingOutsideScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`echo x > doc/guide/a.md`))
}

func TestPolicyRefusesUnknownVariable(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`rm -rf "$HOME/x"`))
}

func TestPolicyForgetsVariableSetInsideLoop(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(
		t,
		p.Allows(`S=tmp/work/a; for d in doc; do S=$d; done; rm -rf "$S"`),
	)
}

func TestPolicyKeepsSubshellChangeDirectoryInside(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`(cd tmp/work) && rm -f x.md`))
}

func TestPolicyRefusesFunctionShadowingAllowlistedName(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`grep() { rm -rf doc; }; grep x a.md`))
}

func TestPolicyRefusesPathOverride(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`PATH=/tmp/evil grep x a.md`))
}

func TestPolicyRefusesFindExec(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`find . -name x -exec rm {} \;`))
}

func TestPolicyAllowsSedInPlaceLikeAllowlist(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`gsed -i 's/a/b/' doc/guide/a.md`))
}

func TestPolicyRefusesUnknownLeadingAssignment(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`FOO=bar app run --rm alpha x`))
}

func TestPolicyAllowsLocaleBeforeCommand(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`LC_ALL=C grep -c x doc/guide/a.md`))
}

func TestPolicyRefusesFieldSeparatorBeforeAnythingButRead(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`IFS=/ ls doc`))
}

func TestPolicyRefusesUnquotedGlobOnWriteCapableCommand(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`find doc -name *.md`))
}

func TestPolicyAllowsQuotedGlobOnWriteCapableCommand(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`find doc -name '*.md'`))
}

func TestPolicyRefusesGitAfterChangingDirectory(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`cd tmp/work && git diff`))
}

func TestPolicyAllowsGitAfterChangeDirectoryThatStays(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(fmt.Sprintf(`cd %s && git diff`, p.Root)))
}

func TestPolicyUnwrapsBareXargs(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`ls doc/guide | xargs grep -c x`))
}

func TestPolicyRefusesXargsWithFlags(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`ls doc/guide | xargs -n1 grep -c x`))
}

func TestPolicyRefusesCommandLookup(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`command -v python3`))
}

func TestPolicyAllowsReadOnlyCommandOutsideAllowlist(t *testing.T) {
	p := policy_tester.New(t)
	assert.True(t, p.Allows(`diff doc/guide/a.md tmp/work/a.md; stat doc`))
}

func TestPolicyRefusesUnparseableCommand(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`for x in; do`))
}

func TestPolicyRefusesBackgroundJob(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`grep x a.md &`))
}

func TestPolicyRefusesMovingOutOfScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`mv doc/guide/a.md tmp/work/`))
}

func TestPolicyRefusesCopyingOutOfScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.False(t, p.Allows(`cp tmp/work/a.md doc/guide/`))
}

func TestPolicyNamesNoBlockerForAllowedCommand(t *testing.T) {
	p := policy_tester.New(t)
	assert.Strings(t, nil, p.Blockers(`ls doc && rm -f tmp/work/x.md`))
}

func TestPolicyNamesEveryBlockerNotOnlyFirst(t *testing.T) {
	p := policy_tester.New(t)
	assert.Strings(
		t,
		[]string{"python3", "rm", "md5"},
		p.Blockers(
			`python3 -c 'print(1)' && rm doc/guide/a.md; python3 -V | md5`,
		),
	)
}

func TestPolicyNamesBlockerInsideLoopsAndSubstitutions(t *testing.T) {
	p := policy_tester.New(t)
	assert.Strings(
		t,
		[]string{"cmp"},
		p.Blockers(`for f in a b; do cmp "$f" x; done`),
	)
	assert.Strings(
		t,
		[]string{"md5"},
		p.Blockers(`echo "$(md5 doc/guide/a.md)"`),
	)
	assert.Strings(
		t,
		[]string{"shasum"},
		p.Blockers(`if ls doc; then shasum doc/guide/a.md; fi`),
	)
}

func TestPolicyWalksOnPastRefusal(t *testing.T) {
	p := policy_tester.New(t)
	both := []string{"cmp", "shasum"}
	assert.Strings(t, both, p.Blockers(`if cmp a b; then shasum x; fi`))
	assert.Strings(t, both, p.Blockers(`while cmp a b; do shasum x; done`))
	assert.Strings(t, both, p.Blockers(`cmp a b || shasum x`))
	assert.Strings(t, both, p.Blockers(`echo "$(cmp a b)$(shasum x)"`))
	assert.Strings(
		t,
		both,
		p.Blockers(`for f in $(cmp a b); do shasum x; done`),
	)
	assert.Strings(t, both, p.Blockers(`case $(cmp a b) in a) shasum x;; esac`))
	assert.Strings(
		t,
		[]string{"(assignment before a command)", "cmp"},
		p.Blockers(`X=1 cmp a b`),
	)
	assert.Strings(
		t,
		[]string{"(assignment to PATH)", "cmp"},
		p.Blockers(`export PATH=/opt HOME=$(cmp a b)`),
	)
	assert.Strings(
		t,
		[]string{"(background job)", "cmp"},
		p.Blockers(`cmp a b &`),
	)
	assert.Strings(
		t,
		[]string{"cmp", "(redirect outside scratch)"},
		p.Blockers(`ls > "$(cmp a b)"`),
	)
	assert.Strings(
		t,
		[]string{"(redirect outside scratch)", "cmp"},
		p.Blockers(`cmp a b > doc/list.txt`),
	)
}

func TestPolicyNamesSubcommandWhereRulesScopeOne(t *testing.T) {
	p := policy_tester.New(t)
	assert.Strings(t, []string{"git status"}, p.Blockers(`git status --short`))
	assert.Strings(t, []string{"app ps"}, p.Blockers(`app ps`))
	assert.Strings(t, []string{"cmp"}, p.Blockers(`cmp a b`))
}

func TestPolicyNamesGitAfterChangeDirectory(t *testing.T) {
	p := policy_tester.New(t)
	assert.Strings(
		t,
		[]string{"(git after cd)"},
		p.Blockers(`cd doc && git diff`),
	)
}

func TestPolicyNamesWriteOutsideScratch(t *testing.T) {
	p := policy_tester.New(t)
	assert.Strings(t, []string{"rm"}, p.Blockers(`rm doc/guide/a.md`))
	assert.Strings(t, []string{"cp"}, p.Blockers(`cp tmp/work/a.md doc/guide/`))
	assert.Strings(
		t,
		[]string{"(redirect outside scratch)"},
		p.Blockers(`ls doc > doc/list.txt`),
	)
}

func TestPolicyNamesStructuralBlocks(t *testing.T) {
	p := policy_tester.New(t)
	assert.Strings(t, []string{"(background job)"}, p.Blockers(`ls doc &`))
	assert.Strings(
		t,
		[]string{"(assignment before a command)"},
		p.Blockers(`X=1 ls doc`),
	)
	assert.Strings(
		t,
		[]string{"(assignment to PATH)"},
		p.Blockers(`export PATH=/opt/bin; ls doc`),
	)
	assert.Strings(
		t,
		[]string{"(array assignment)"},
		p.Blockers(`a=(1 2); ls doc`),
	)
	assert.Strings(
		t,
		[]string{"(command from an expansion)"},
		p.Blockers(`$tool doc`),
	)
	assert.Strings(
		t,
		[]string{"(unsupported construct)"},
		p.Blockers(`f() { ls doc; }`),
	)
	assert.Strings(t, []string{"(unparsed)"}, p.Blockers(`ls doc (`))
	assert.Strings(t, []string{"(no command)"}, p.Blockers(`x=1`))
	assert.Strings(t, []string{"find"}, p.Blockers(`find doc -delete`))
}

func TestPolicyNamesWrapperItCannotStrip(t *testing.T) {
	p := policy_tester.New(t)
	assert.Strings(t, []string{"xargs"}, p.Blockers(`ls doc | xargs -n1 cat`))
	assert.Strings(t, []string{"timeout"}, p.Blockers(`timeout 5`))
	assert.Strings(t, []string{"cmp"}, p.Blockers(`timeout 5 cmp a b`))
}
