package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/permission/scratch"
	"github.com/funtimecoding/soil/pkg/generative/unit/policy_tester"
	"path/filepath"
	"testing"
)

func TestPolicyRedirectUnderScratchPrefixAllowed(t *testing.T) {
	root, scratchpad := policy_tester.Project(t)
	prefix := filepath.Join(t.TempDir(), "alfa-")
	assert.True(
		t,
		policy_tester.AllowsWithPrefix(
			t,
			root,
			scratchpad,
			prefix,
			fmt.Sprintf("grep -c x doc/guide/a.md > %sreport.txt", prefix),
		),
	)
	assert.True(
		t,
		policy_tester.AllowsWithPrefix(
			t,
			root,
			scratchpad,
			prefix,
			fmt.Sprintf("grep -c x doc/guide/a.md > %swork/count.txt", prefix),
		),
	)
}

func TestPolicyRedirectBesideScratchPrefixBlocked(t *testing.T) {
	root, scratchpad := policy_tester.Project(t)
	directory := t.TempDir()
	prefix := filepath.Join(directory, "alfa-")

	for _, target := range []string{
		filepath.Join(directory, "alfabet.txt"),
		filepath.Join(directory, "bravo-report.txt"),
		filepath.Join(
			directory,
			"alfa-work",
			soil.ParentDirectory,
			soil.ParentDirectory,
			"report.txt",
		),
	} {
		assert.False(
			t,
			policy_tester.AllowsWithPrefix(
				t,
				root,
				scratchpad,
				prefix,
				fmt.Sprintf("grep -c x doc/guide/a.md > %s", target),
			),
		)
	}
}

func TestPolicyRedirectUnderPrefixBlockedWithoutIt(t *testing.T) {
	root, scratchpad := policy_tester.Project(t)
	prefix := filepath.Join(t.TempDir(), "alfa-")
	assert.False(
		t,
		policy_tester.Allows(
			t,
			root,
			scratchpad,
			fmt.Sprintf("grep -c x doc/guide/a.md > %sreport.txt", prefix),
		),
	)
}

func TestScratchPrefixLowercased(t *testing.T) {
	assert.String(t, "/tmp/alfa-", scratch.Prefix("Alfa"))
}

func TestScratchRootsIncludeEachDirectory(t *testing.T) {
	assert.Strings(
		t,
		[]string{"/srv/alfa/tmp", "/srv/bravo/tmp", "/srv/charlie/tmp"},
		scratch.Roots("/srv/alfa", []string{"/srv/bravo", "/srv/charlie"}),
	)
}
