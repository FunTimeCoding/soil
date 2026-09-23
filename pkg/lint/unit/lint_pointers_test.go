package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer_tester"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestPointersGitignored(t *testing.T) {
	l := pointer_tester.Gitignored(".claude/notes/alfa.md")(
		constant.UpperAlfa,
		strings.NewReader("Style guide at `.claude/notes/alfa.md` today.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersClean(t *testing.T) {
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader("Run `task lint` now.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersExisting(t *testing.T) {
	l := pointer_tester.Checker("doc/ai/runbook/lint.md")(
		constant.UpperAlfa,
		strings.NewReader("Read `doc/ai/runbook/lint.md` first.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersDirectory(t *testing.T) {
	l := pointer_tester.Checker("doc/ai/spec")(
		constant.UpperAlfa,
		strings.NewReader("Specs live in `doc/ai/spec/`.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersIgnored(t *testing.T) {
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(
			"See `tmp/gosec.json`, `doc/ai/runbook/<name>.md`, and https://example.org/page.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersDead(t *testing.T) {
	line := "Read `doc/ai/runbook/ghost.md` first."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersAbsolute(t *testing.T) {
	line := "Notes at `/Users/example/notes.md` today."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Absolute("Alfa", line), "", l)
}

func TestPointersRelativeExisting(t *testing.T) {
	l := pointer_tester.Checker("doc/ai/spec/naming.md")(
		"doc/ai/runbook/lint.md",
		strings.NewReader("See `../spec/naming.md` for the rules.\n"),
	)
	assertReport(t, "doc/ai/runbook/lint.md", false, nil, "", l)
}

func TestPointersRelativeDead(t *testing.T) {
	line := "See `../spec/ghost.md` for the rules."
	l := pointer_tester.Checker("doc/ai/spec/naming.md")(
		"doc/ai/runbook/lint.md",
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"doc/ai/runbook/lint.md",
		true,
		pointer_tester.Dead("doc/ai/runbook/lint.md", line, 1),
		"",
		l,
	)
}

func TestPointersSymbolIgnored(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision/salt", "pkg/system/run")(
		constant.UpperAlfa,
		strings.NewReader(
			"Wraps `go:pkg/provision/salt.Client` and `go:pkg/system/run/New()`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersSymbolDeadPackage(t *testing.T) {
	line := "Wraps `go:pkg/provision/salt.Client` today."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersSymbolStdlib(t *testing.T) {
	l := pointer_tester.Stdlib("crypto/x509", "go/packages")(
		constant.UpperAlfa,
		strings.NewReader(
			"Parses via `go:crypto/x509` and loads with `go:go/packages.Load`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersSymbolStdlibSingleSegment(t *testing.T) {
	l := pointer_tester.Stdlib("fmt", "strings")(
		constant.UpperAlfa,
		strings.NewReader(
			"Prints with `go:fmt.Println` and cuts with `go:strings`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersSymbolSingleSegmentDead(t *testing.T) {
	line := "Prints with `go:fmt.Println` today."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersSymbolWithoutPackageDead(t *testing.T) {
	line := "Calls `go:Whatever` today."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersUncheckedEnumerates(t *testing.T) {
	checker, seen := pointer_tester.Recording()
	l := checker(
		constant.UpperAlfa,
		strings.NewReader(
			"Serves `route:/health`, reads `etc/hosts`, fetches `example.org/module`, keeps `~/.config/example`, allows `10.0.0.0/24` and mounts `192.0.2.8:/export`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
	assert.Strings(
		t,
		[]string{
			"Alfa:1 route route:/health",
			"Alfa:1 unknown etc/hosts",
			"Alfa:1 external example.org/module",
			"Alfa:1 home ~/.config/example",
			"Alfa:1 network 10.0.0.0/24",
			"Alfa:1 network 192.0.2.8:/export",
		},
		*seen,
	)
}

func TestPointersConfiguredRegistry(t *testing.T) {
	checker, seen := pointer_tester.Registries("registry.example")
	l := checker(
		constant.UpperAlfa,
		strings.NewReader(
			"Pulls `registry.example/team/image` and `other.example/team/image`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
	assert.Strings(t, []string{"image", "external"}, *seen)
}

func TestPointersConventionWord(t *testing.T) {
	line := "Constants live in `constant/constant.go` and tests in `unit/`."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"Alfa",
		true,
		pointer_tester.Convention("Alfa", line, 2),
		"",
		l,
	)
}

func TestPointersRelativeLinkWithDirectory(t *testing.T) {
	l := pointer_tester.Checker("doc/guide/client/alpha.md")(
		"doc/guide/clients.md",
		strings.NewReader(
			"| Alpha | JSON API | [client/alpha.md](client/alpha.md) |\n",
		),
	)
	assertReport(t, "doc/guide/clients.md", false, nil, "", l)
}

func TestPointersAncestorResolves(t *testing.T) {
	l := pointer_tester.Checker(
		"doc/guide/alpha/charlie",
		"doc/guide/alpha/charlie/Delta.md",
	)(
		"doc/guide/alpha/reader/README.md",
		strings.NewReader(
			"Charlie lives in `charlie/` and delta in `charlie/Delta.md`.\n",
		),
	)
	assertReport(t, "doc/guide/alpha/reader/README.md", false, nil, "", l)
}

func TestPointersAncestorPrefixResolves(t *testing.T) {
	l := pointer_tester.Checker(
		"doc/notes/alpha",
		"doc/notes/alpha/a8-foxtrot.md",
	)(
		"doc/notes/bravo/b303-echo.md",
		strings.NewReader("Continues `alpha/a8`.\n"),
	)
	assertReport(t, "doc/notes/bravo/b303-echo.md", false, nil, "", l)
}

func TestPointersAncestorPrefixMissingTallies(t *testing.T) {
	checker, count := pointer_tester.Counting("doc/notes/alpha")
	l := checker(
		"doc/notes/bravo/b303-echo.md",
		strings.NewReader("Continues `alpha/a9`.\n"),
	)
	assertReport(t, "doc/notes/bravo/b303-echo.md", false, nil, "", l)
	assert.Integer(t, 1, *count)
}

func TestPointersDependencyLive(t *testing.T) {
	l := pointer_tester.Dependencies(
		"github.com/ory/fosite",
		"gorm.io/driver/sqlite",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			"Uses `ory/fosite`, `ory/fosite/handler`, `gorm.io/driver/sqlite` and `github.com/ory/fosite`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersImplicitBaseLive(t *testing.T) {
	l := pointer_tester.Implicit(
		[]string{lintConstant.PackageDirectory, "../github/soil/pkg"},
		lintConstant.PackageDirectory,
		"../github/soil/pkg",
		"../github/soil/pkg/generative",
		"../github/soil/pkg/generative/constant",
	)(
		constant.UpperAlfa,
		strings.NewReader("Vocabulary in `generative/constant`.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersImplicitBaseDead(t *testing.T) {
	line := "Vocabulary in `generative/gone` today."
	l := pointer_tester.Implicit(
		[]string{lintConstant.PackageDirectory, "../github/soil/pkg"},
		lintConstant.PackageDirectory,
		"../github/soil/pkg",
		"../github/soil/pkg/generative",
	)(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersBracePlaceholderIgnored(t *testing.T) {
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(
			"Files land in `{zooPath}/vid/` and `{world}/{filename}`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersConventionWordAnchored(t *testing.T) {
	l := pointer_tester.Checker("pkg/foo", "pkg/foo/constant")(
		constant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/foo\n---\nConstants live in `constant/` here.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersBasedSeveral(t *testing.T) {
	l := pointer_tester.Checker(
		"../github/soil/pkg/tool/goalertlogd",
		"../github/soil/pkg/tool/goalertlogd/store",
		"container",
		"container/alpha",
		"container/alpha/compose.yaml",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			"---\nbase: ../github/soil/pkg/tool/goalertlogd, container\n---\nRecords sit in `store/` and the service in `alpha/compose.yaml`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersBasedSeveralInteriorDead(t *testing.T) {
	line := "The service is in `alpha/compose.yml` today."
	l := pointer_tester.Checker(
		"pkg/tool/goalertlogd",
		"container",
		"container/alpha",
		"container/alpha/compose.yaml",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf(
				"---\nbase: pkg/tool/goalertlogd, container\n---\n%s\n",
				line,
			),
		),
	)
	assertReport(t, "Alfa", true, pointer_tester.DeadAt("Alfa", 4, line), "", l)
}

func TestPointersUnprefixedSymbolIsAPath(t *testing.T) {
	line := "Wraps `pkg/provision/salt.Client` today."
	l := pointer_tester.Checker("pkg/provision/salt")(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersSymbolPlaceholderIgnored(t *testing.T) {
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader("Write `go:pkg/<name>/Symbol` in examples.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersUppercaseFileDead(t *testing.T) {
	line := "Build from `container/gopostgres/Containerfile` here."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersUppercaseFileExisting(t *testing.T) {
	l := pointer_tester.Checker("container/gopostgres/Containerfile")(
		constant.UpperAlfa,
		strings.NewReader(
			"Build from `container/gopostgres/Containerfile` here.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersSiblingSymbolIgnored(t *testing.T) {
	l := pointer_tester.Checker("../github/soil/pkg/provision/salt")(
		constant.UpperAlfa,
		strings.NewReader(
			"Wraps `go:../github/soil/pkg/provision/salt.Client` today.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersSiblingSymbolDeadPackage(t *testing.T) {
	line := "Wraps `go:../github/soil/pkg/gone/Thing` today."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersReceiverMethodIgnored(t *testing.T) {
	l := pointer_tester.Checker("pkg/lint/pointer")(
		constant.UpperAlfa,
		strings.NewReader("Calls `go:pkg/lint/pointer/Names.ResolvePackage`.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersQualifiedMethodIgnored(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision/salt")(
		constant.UpperAlfa,
		strings.NewReader("Calls `go:pkg/provision/salt.Client.Login`.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersSymbolExpansion(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision/salt", "pkg/provision/ansible")(
		constant.UpperAlfa,
		strings.NewReader(
			"Both `go:pkg/provision/{salt,ansible}.Client` work.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersLiveSymbolInLivePackage(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision/salt")(
		constant.UpperAlfa,
		strings.NewReader("Renamed away: `go:pkg/provision/salt.Gone`.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersBasedAnchorResolves(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision", "pkg/provision/salt")(
		constant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/provision\n---\nConnectors live in `salt/` today.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersBasedInteriorDead(t *testing.T) {
	line := "Connectors live in `salt/gone.go` today."
	l := pointer_tester.Checker("pkg/provision", "pkg/provision/salt")(
		constant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nbase: pkg/provision\n---\n%s\n", line),
		),
	)
	assertReport(t, "Alfa", true, pointer_tester.DeadAt("Alfa", 4, line), "", l)
}

func TestPointersBasedForeignNamespaceIgnored(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision")(
		constant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/provision\n---\nGroups sit in `keepass/General` outside.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersBasedSibling(t *testing.T) {
	l := pointer_tester.Checker(
		"../github/soil/pkg/tool/gomaintlogd",
		"../github/soil/pkg/tool/gomaintlogd/server",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			"---\nbase: ../github/soil/pkg/tool/gomaintlogd\n---\nHandlers sit in `server/` today.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersBasedSiblingFile(t *testing.T) {
	l := pointer_tester.Checker(
		"../github/soil/pkg/tool/gomaintlogd",
		"../github/soil/pkg/tool/gomaintlogd/face",
		"../github/soil/pkg/tool/gomaintlogd/face/source.go",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			"---\nbase: ../github/soil/pkg/tool/gomaintlogd\n---\nThe reader lives in `face/source.go` today.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersBasedSiblingInteriorDead(t *testing.T) {
	line := "Handlers sit in `server/gone.go` today."
	l := pointer_tester.Checker(
		"../github/soil/pkg/tool/gomaintlogd",
		"../github/soil/pkg/tool/gomaintlogd/server",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf(
				"---\nbase: ../github/soil/pkg/tool/gomaintlogd\n---\n%s\n",
				line,
			),
		),
	)
	assertReport(t, "Alfa", true, pointer_tester.DeadAt("Alfa", 4, line), "", l)
}

func TestPointersCommandResolves(t *testing.T) {
	checker, count := pointer_tester.Counting(".claude/skills/sign-firefox")
	l := checker(
		constant.UpperAlfa,
		strings.NewReader("Invoke `/sign-firefox` before pushing.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersCommandPluginResolves(t *testing.T) {
	checker, count := pointer_tester.Counting(
		"../soil/strata/plugin/soil/skills/lint",
	)
	l := checker(
		constant.UpperAlfa,
		strings.NewReader("Invoke `/soil:lint` after the scope.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersCommandDead(t *testing.T) {
	line := "Invoke `/ghost-skill` before pushing."
	checker, count := pointer_tester.Counting()
	l := checker(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersCommandHarness(t *testing.T) {
	checker, count := pointer_tester.Counting()
	l := checker(
		constant.UpperAlfa,
		strings.NewReader("New tools need a `/mcp` reconnect.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersCommandPluginInRepository(t *testing.T) {
	checker, count := pointer_tester.Counting("strata/plugin/soil/skills/lint")
	l := checker(
		constant.UpperAlfa,
		strings.NewReader("Invoke `/soil:lint` after the scope.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersRouteTallies(t *testing.T) {
	checker, count := pointer_tester.Counting()
	l := checker(
		constant.UpperAlfa,
		strings.NewReader(
			"The service answers on `route:/health` for probes.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
	assert.Integer(t, 1, *count)
}

func TestPointersRouteWithoutSpecificationDead(t *testing.T) {
	line := "Lists at `route:/api/alerts`."
	l := pointer_tester.Checker("pkg/tool/alpha")(
		constant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nbase: pkg/tool/alpha\n---\n%s\n", line),
		),
	)
	assertReport(t, "Alfa", true, pointer_tester.DeadAt("Alfa", 4, line), "", l)
}

func TestPointersRouteLive(t *testing.T) {
	l := pointer_tester.Routes(
		map[string][]string{
			"pkg/tool/alpha": {"/api/alerts", "/api/alerts/{name}"},
		},
		"pkg/tool/alpha",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/tool/alpha\n---\nLists at `route:/api/alerts`, one at `route:/api/alerts/{identifier}`, filtered by `route:/api/alerts?name=x` and `route:/api/...`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersRouteDead(t *testing.T) {
	line := "Lists at `route:/api/ghosts` today."
	l := pointer_tester.Routes(
		map[string][]string{"pkg/tool/alpha": {"/api/alerts"}},
		"pkg/tool/alpha",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nbase: pkg/tool/alpha\n---\n%s\n", line),
		),
	)
	assertReport(t, "Alfa", true, pointer_tester.DeadAt("Alfa", 4, line), "", l)
}

func TestPointersRouteLiteralLive(t *testing.T) {
	l := pointer_tester.Literals(
		map[string][]string{
			"pkg/tool/alpha": {
				`WidgetsPath = "/widgets"`,
				`mux.HandleFunc("/memories/", list)`,
			},
		},
		nil,
		nil,
		"pkg/tool/alpha",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/tool/alpha\n---\nPages at `route:/widgets` and `route:/memories/{identifier}`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersRouteLiteralDead(t *testing.T) {
	line := "Pages at `route:/ghosts` today."
	l := pointer_tester.Literals(
		map[string][]string{"pkg/tool/alpha": {`WidgetsPath = "/widgets"`}},
		nil,
		nil,
		"pkg/tool/alpha",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nbase: pkg/tool/alpha\n---\n%s\n", line),
		),
	)
	assertReport(t, "Alfa", true, pointer_tester.DeadAt("Alfa", 4, line), "", l)
}

func TestPointersRouteLiteralInSharedTree(t *testing.T) {
	l := pointer_tester.Literals(
		map[string][]string{
			"pkg/tool/alpha":       {`WidgetsPath = "/widgets"`},
			"../github/shared/pkg": {`HealthPath = "/health"`},
		},
		nil,
		[]string{"../github/shared/pkg"},
		"pkg/tool/alpha",
	)(
		constant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/tool/alpha\n---\nProbes `route:/health`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersImageLive(t *testing.T) {
	l := pointer_tester.Literals(
		map[string][]string{"": {"image: ghcr.io/example/alpha:v1"}},
		nil,
		nil,
	)(
		constant.UpperAlfa,
		strings.NewReader("Runs `ghcr.io/example/alpha`.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersBareSlashDead(t *testing.T) {
	line := "Lists at `/api/alerts` and `/api/alerts/{name}`."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"Alfa",
		true,
		pointer_tester.BareSlash("Alfa", line, 2),
		"",
		l,
	)
}

func TestPointersSystemTallies(t *testing.T) {
	checker, seen := pointer_tester.Recording()
	l := checker(
		constant.UpperAlfa,
		strings.NewReader(
			"Mounts `/proc/`, reads `/etc/app.conf`, runs `path:/bin/true`, forgets `/bin/false`, pulls `ghcr.io/example/image`, imports `\"github.com/example/module\"`, rewrites with `s/a/b/`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
	assert.Strings(
		t,
		[]string{
			"Alfa:1 system /proc/",
			"Alfa:1 system /etc/app.conf",
			"Alfa:1 system path:/bin/true",
			"Alfa:1 slash /bin/false",
			"Alfa:1 image ghcr.io/example/image",
			"Alfa:1 import \"github.com/example/module\"",
			"Alfa:1 pattern s/a/b/",
		},
		*seen,
	)
}

func TestPointersBaseDead(t *testing.T) {
	line := "base: pkg/gone"
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("---\n%s\n---\nProse.\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.DeadAt("Alfa", 2, line), "", l)
}

func TestPointersExpansionExisting(t *testing.T) {
	l := pointer_tester.Checker(
		"doc/ai/spec/naming.md",
		"doc/ai/spec/build.md",
	)(
		constant.UpperAlfa,
		strings.NewReader("Read `doc/ai/spec/{naming,build}.md` first.\n"),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersExpansionDead(t *testing.T) {
	line := "Read `doc/ai/spec/{naming,ghost}.md` first."
	l := pointer_tester.Checker("doc/ai/spec/naming.md")(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersSiblingExisting(t *testing.T) {
	l := pointer_tester.Checker("../../github/soil/doc/ai/spec/naming.md")(
		constant.UpperAlfa,
		strings.NewReader(
			"Shared specs at `../../github/soil/doc/ai/spec/naming.md`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersSiblingDead(t *testing.T) {
	line := "Shared specs at `../github/soil/doc/ai/spec/naming.md`."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 1), "", l)
}

func TestPointersMultiple(t *testing.T) {
	line := "Read `doc/ai/spec/ghost.md` and `pkg/gone`."
	l := pointer_tester.Checker()(
		constant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(t, "Alfa", true, pointer_tester.Dead("Alfa", line, 2), "", l)
}

func TestPointersPluginRoot(t *testing.T) {
	l := pointer_tester.Checker("doc/ai/runbook/lint.md")(
		constant.UpperAlfa,
		strings.NewReader(
			"Read `${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md`.\n",
		),
	)
	assertReport(t, "Alfa", false, nil, "", l)
}

func TestPointersBareLinkDead(t *testing.T) {
	line := "See [ghost](ghost.md) for details."
	l := pointer_tester.Checker()(
		"doc/guide/index.md",
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	assertReport(
		t,
		"doc/guide/index.md",
		true,
		pointer_tester.Dead("doc/guide/index.md", line, 1),
		"",
		l,
	)
}

func TestPointersBareLinkExisting(t *testing.T) {
	l := pointer_tester.Checker("doc/guide/other.md")(
		"doc/guide/index.md",
		strings.NewReader("See [other](other.md) for details.\n"),
	)
	assertReport(t, "doc/guide/index.md", false, nil, "", l)
}

func TestPointersBareLinkAnchorIgnored(t *testing.T) {
	l := pointer_tester.Checker()(
		"doc/guide/index.md",
		strings.NewReader("Jump to [section](#section) below.\n"),
	)
	assertReport(t, "doc/guide/index.md", false, nil, "", l)
}

func TestPointersBareLinkAnchorSuffix(t *testing.T) {
	l := pointer_tester.Checker("doc/guide/other.md")(
		"doc/guide/index.md",
		strings.NewReader("See [other](other.md#part) there.\n"),
	)
	assertReport(t, "doc/guide/index.md", false, nil, "", l)
}
