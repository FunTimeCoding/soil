package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/guard"
	"testing"
)

func TestVerdictBlocksNpx(t *testing.T) {
	assert.String(
		t,
		"npx is blocked (supply-chain guard) - it downloads and executes npm packages on demand",
		guard.Verdict("darwin", "npx cowsay hi"),
	)
	assert.StringContains(
		t,
		"npx is blocked",
		guard.Verdict("darwin", "cd web && npx tsc --noEmit"),
	)
	assert.StringContains(
		t,
		"npx is blocked",
		guard.Verdict("darwin", "/usr/local/bin/npx serve ."),
	)
	assert.StringContains(
		t,
		"npx is blocked",
		guard.Verdict("linux", "npx create-react-app demo"),
	)
}

func TestVerdictBlocksPipInstall(t *testing.T) {
	assert.String(
		t,
		"pip install is blocked (supply-chain guard) - no python dependencies may be installed on this system",
		guard.Verdict("darwin", "pip install requests"),
	)
	assert.StringContains(
		t,
		"pip install is blocked",
		guard.Verdict("darwin", "pip3 install --upgrade requests"),
	)
	assert.StringContains(
		t,
		"pip install is blocked",
		guard.Verdict("darwin", "python -m pip install requests"),
	)
	assert.StringContains(
		t,
		"pip install is blocked",
		guard.Verdict("linux", "python3 -m pip install -U requests"),
	)
	assert.StringContains(
		t,
		"pip install is blocked",
		guard.Verdict("darwin", "pip --no-cache-dir install requests"),
	)
}

func TestVerdictAllowsPackageQueries(t *testing.T) {
	assert.String(t, "", guard.Verdict("darwin", "pip list"))
	assert.String(t, "", guard.Verdict("darwin", "pip3 show requests"))
	assert.String(t, "", guard.Verdict("darwin", "python -m venv .venv"))
	assert.String(t, "", guard.Verdict("darwin", "npm ls"))
	assert.String(
		t,
		"",
		guard.Verdict("darwin", `ssh host.example "npx cowsay hi"`),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", `ssh host.example "pip install requests"`),
	)
}

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
	assert.String(t, "", guard.Verdict("darwin", "grep x file | sed 's/a/b/'"))
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
		guard.Verdict("darwin", `grep -cE '(^|[|&;( ])sed( |$)' corpus.txt`),
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

func TestVerdictBlocksLocalXargsUnsupportedFlag(t *testing.T) {
	assert.String(
		t,
		"xargs on macOS is BSD xargs and rejects the GNU flags -a and -d - redirect the file on stdin instead: xargs command < file",
		guard.Verdict("darwin", "xargs -a list.txt git checkout --"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "xargs --arg-file=list.txt rm"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "xargs -d, echo"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "xargs --delimiter=, echo"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "/usr/bin/xargs -a list.txt echo"),
	)
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "cat other && xargs -a list.txt echo"),
	)
}

func TestVerdictAllowsRemoteXargs(t *testing.T) {
	assert.String(
		t,
		"",
		guard.Verdict("darwin", `ssh root@10.0.0.2 "xargs -a list.txt rm"`),
	)
	assert.String(
		t,
		"",
		guard.Verdict(
			"darwin",
			`podman run --rm alpine sh -c "xargs -a list.txt echo"`,
		),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", `kubectl exec pod -- sh -c "xargs -d, echo"`),
	)
}

func TestVerdictAllowsSupportedXargsFlags(t *testing.T) {
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "xargs git checkout -- < list.txt"),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "find . -print0 | xargs -0 rm"),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "grep -l x * | xargs -I{} echo {}"),
	)
	assert.String(
		t,
		"",
		guard.Verdict("darwin", "cat list | xargs -n1 -P4 echo"),
	)
	assert.String(t, "", guard.Verdict("darwin", "cat list | xargs -r echo"))
	assert.String(t, "", guard.Verdict("darwin", "cat list | xargs -L1 echo"))
}

func TestVerdictAllowsXargsOnLinux(t *testing.T) {
	assert.String(t, "", guard.Verdict("linux", "xargs -a list.txt echo"))
}

func TestVerdictXargsParseFallback(t *testing.T) {
	assert.StringContains(
		t,
		"redirect the file on stdin",
		guard.Verdict("darwin", "xargs -a 'list.txt echo"),
	)
}
