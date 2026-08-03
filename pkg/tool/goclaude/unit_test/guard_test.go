package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestVerdictBlocksLocalInPlace(t *testing.T) {
	assert.String(
		t,
		"sed on macOS is BSD sed and its flags (notably -i) differ from GNU sed - use gsed instead",
		guard.Verdict("darwin", "sed -i 's/a/b/' file.go"),
	)
	assert.StringContains(
		t,
		"use gsed instead",
		guard.Verdict("darwin", "sed -i.bak 's/a/b/' file.go"),
	)
	assert.StringContains(
		t,
		"use gsed instead",
		guard.Verdict("darwin", "sed --in-place 's/a/b/' file.go"),
	)
	assert.StringContains(
		t,
		"use gsed instead",
		guard.Verdict("darwin", "/usr/bin/sed -i 's/a/b/' file.go"),
	)
	assert.StringContains(
		t,
		"use gsed instead",
		guard.Verdict("darwin", "LC_ALL=C sed -i 's/a/b/' file.go"),
	)
	assert.StringContains(
		t,
		"use gsed instead",
		guard.Verdict(
			"darwin",
			"ssh host.example 'cat remote.txt' | sed -i 's/a/b/' local.txt",
		),
	)
	assert.StringContains(
		t,
		"use gsed instead",
		guard.Verdict("darwin", "cat file && sed -i 's/a/b/' file"),
	)
}

func TestVerdictAllowsRemote(t *testing.T) {
	assert.String(
		t,
		"",
		guard.Verdict(
			"darwin",
			`ssh root@10.0.0.2 "sed -i 's/a/b/' /etc/app.conf"`,
		),
	)
	assert.String(
		t,
		"",
		guard.Verdict(
			"darwin",
			`ssh root@host.example "sh -c 'sed -i s/a/b/ /tmp/file'"`,
		),
	)
	assert.String(
		t,
		"",
		guard.Verdict(
			"darwin",
			`ssh admin@server.test "sudo sed -i 's/a/b/' /etc/app.conf"`,
		),
	)
}

func TestVerdictAllowsFilter(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", "sed 's/a/b/' file.go"))
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "grep x file | sed 's/a/b/'"),
	)
	assert.String(t, "", guard.Verdict("darwin", "cat file; sed -n 1p file"))
	assert.String(
		t,
		"",
		guard.Verdict(
			"darwin",
			`ssh host.example './tool' | sed 's|prefix: ||'`,
		),
	)
}

func TestVerdictAllowsSedAsArgument(t *testing.T) {
	assert.String(
		t,
		"",
		guard.Verdict(
			"darwin",
			`grep -cE '(^|[|&;( ])sed( |$)' corpus.txt`,
		),
	)
	assert.String(t, "", guard.Verdict("darwin", "echo parsed"))
	assert.String(t, "", guard.Verdict("darwin", "git grep sedative"))
}

func TestVerdictAllowsOther(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", "gsed -i 's/a/b/' file.go"))
	assert.String(t, "", guard.Verdict("linux", "sed -i 's/a/b/' file.go"))
}

func TestVerdictParseFallback(t *testing.T) {
	assert.StringContains(
		t,
		"use gsed instead",
		guard.Verdict("darwin", "sed -i 's/a/b/ file.go"),
	)
	assert.String(t, "", guard.Verdict("darwin", "echo 'unclosed"))
}
