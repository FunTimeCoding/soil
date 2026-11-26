package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/pointer_tester"
	"github.com/funtimecoding/soil/pkg/lint/unit/lint_tester"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
	"testing"
)

func TestPointersGitignored(t *testing.T) {
	l := pointer_tester.Gitignored(".claude/notes/alfa.md")(
		stringsConstant.UpperAlfa,
		strings.NewReader("Style guide at `.claude/notes/alfa.md` today.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersClean(t *testing.T) {
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader("Run `task lint` now.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersExisting(t *testing.T) {
	l := pointer_tester.Checker("doc/ai/runbook/lint.md")(
		stringsConstant.UpperAlfa,
		strings.NewReader("Read `doc/ai/runbook/lint.md` first.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersDirectory(t *testing.T) {
	l := pointer_tester.Checker("doc/ai/spec")(
		stringsConstant.UpperAlfa,
		strings.NewReader("Specs live in `doc/ai/spec/`.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersIgnored(t *testing.T) {
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"See `tmp/gosec.json`, `doc/ai/runbook/<name>.md`, and https://example.org/page.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersDead(t *testing.T) {
	line := "Read `doc/ai/runbook/ghost.md` first."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersDeadWithSpace(t *testing.T) {
	line := "Read `doc/guide/Bravo Charlie.md` first."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersExistingWithSpace(t *testing.T) {
	l := pointer_tester.Checker("doc/guide/Bravo Charlie.md")(
		stringsConstant.UpperAlfa,
		strings.NewReader("Read `doc/guide/Bravo Charlie.md` first.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersAbsolute(t *testing.T) {
	line := "Notes at `/Users/example/notes.md` today."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Absolute("Alfa", line),
		"",
		l,
	)
}

func TestPointersRelativeExisting(t *testing.T) {
	l := pointer_tester.Checker("doc/ai/spec/naming.md")(
		"doc/ai/runbook/lint.md",
		strings.NewReader("See `../spec/naming.md` for the rules.\n"),
	)
	lint_tester.New(t).Report("doc/ai/runbook/lint.md", false, nil, "", l)
}

func TestPointersRelativeDead(t *testing.T) {
	line := "See `../spec/ghost.md` for the rules."
	l := pointer_tester.Checker("doc/ai/spec/naming.md")(
		"doc/ai/runbook/lint.md",
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"doc/ai/runbook/lint.md",
		true,
		pointer_tester.Dead("doc/ai/runbook/lint.md", line, 1),
		"",
		l,
	)
}

func TestPointersSymbolIgnored(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision/salt", "pkg/system/run")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Wraps `go:pkg/provision/salt.Client` and `go:pkg/system/run/New()`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersSymbolDeadPackage(t *testing.T) {
	line := "Wraps `go:pkg/provision/salt.Client` today."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersSymbolStdlib(t *testing.T) {
	l := pointer_tester.Stdlib("crypto/x509", "go/packages")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Parses via `go:crypto/x509` and loads with `go:go/packages.Load`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersSymbolStdlibSingleSegment(t *testing.T) {
	l := pointer_tester.Stdlib("fmt", "strings")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Prints with `go:fmt.Println` and cuts with `go:strings`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersSymbolSingleSegmentDead(t *testing.T) {
	line := "Prints with `go:fmt.Println` today."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersSymbolWithoutPackageDead(t *testing.T) {
	line := "Calls `go:Whatever` today."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersUncheckedEnumerates(t *testing.T) {
	checker, seen := pointer_tester.Recording()
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Serves `route:/health`, reads `etc/hosts`, fetches `example.org/module`, keeps `~/.config/example`, allows `10.0.0.0/24` and mounts `192.0.2.8:/export`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
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
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Pulls `registry.example/team/image` and `other.example/team/image`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Strings(t, []string{"image", "external"}, *seen)
}

func TestPointersConventionWord(t *testing.T) {
	line := "Constants live in `constant/constant.go` and tests in `unit/`."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Convention("Alfa", line, 2),
		"",
		l,
	)
}

func TestPointersRelativeLinkWithDirectory(t *testing.T) {
	l := pointer_tester.Checker("doc/guide/client/alfa.md")(
		"doc/guide/clients.md",
		strings.NewReader(
			"| Alfa | JSON API | [client/alfa.md](client/alfa.md) |\n",
		),
	)
	lint_tester.New(t).Report("doc/guide/clients.md", false, nil, "", l)
}

func TestPointersAncestorResolves(t *testing.T) {
	l := pointer_tester.Checker(
		"doc/guide/alfa/charlie",
		"doc/guide/alfa/charlie/Delta.md",
	)(
		"doc/guide/alfa/reader/README.md",
		strings.NewReader(
			"Charlie lives in `charlie/` and delta in `charlie/Delta.md`.\n",
		),
	)
	lint_tester.New(t).Report(
		"doc/guide/alfa/reader/README.md",
		false,
		nil,
		"",
		l,
	)
}

func TestPointersAncestorPrefixResolves(t *testing.T) {
	l := pointer_tester.Checker(
		"doc/notes/alfa",
		"doc/notes/alfa/a8-foxtrot.md",
	)(
		"doc/notes/bravo/b303-echo.md",
		strings.NewReader("Continues `alfa/a8`.\n"),
	)
	lint_tester.New(t).Report("doc/notes/bravo/b303-echo.md", false, nil, "", l)
}

func TestPointersAncestorPrefixMissingUnchecked(t *testing.T) {
	checker, count := pointer_tester.Counting("doc/notes/alfa")
	l := checker(
		"doc/notes/bravo/b303-echo.md",
		strings.NewReader("Continues `alfa/a9`.\n"),
	)
	lint_tester.New(t).Report("doc/notes/bravo/b303-echo.md", false, nil, "", l)
	assert.Integer(t, 1, *count)
}

func TestPointersDependencyLive(t *testing.T) {
	l := pointer_tester.Dependencies(
		"github.com/ory/fosite",
		"gorm.io/driver/sqlite",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Uses `ory/fosite`, `ory/fosite/handler`, `gorm.io/driver/sqlite` and `github.com/ory/fosite`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersImplicitBaseLive(t *testing.T) {
	l := pointer_tester.Implicit(
		[]string{constant.PackageDirectory, "../github/soil/pkg"},
		constant.PackageDirectory,
		"../github/soil/pkg",
		"../github/soil/pkg/generative",
		"../github/soil/pkg/generative/constant",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader("Vocabulary in `generative/constant`.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersImplicitBaseDead(t *testing.T) {
	line := "Vocabulary in `generative/gone` today."
	l := pointer_tester.Implicit(
		[]string{constant.PackageDirectory, "../github/soil/pkg"},
		constant.PackageDirectory,
		"../github/soil/pkg",
		"../github/soil/pkg/generative",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersBracePlaceholderIgnored(t *testing.T) {
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Files land in `{zooPath}/vid/` and `{world}/{filename}`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersConventionWordAnchored(t *testing.T) {
	l := pointer_tester.Checker("pkg/foo", "pkg/foo/constant")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/foo\n---\nConstants live in `constant/` here.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersBasedSeveral(t *testing.T) {
	l := pointer_tester.Checker(
		"../github/soil/pkg/tool/goalertlogd",
		"../github/soil/pkg/tool/goalertlogd/store",
		"strata/container",
		"strata/container/alfa",
		"strata/container/alfa/compose.yaml",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nbase: ../github/soil/pkg/tool/goalertlogd, strata/container\n---\nRecords sit in `store/` and the service in `alfa/compose.yaml`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersBasedSeveralInteriorDead(t *testing.T) {
	line := "The service is in `alfa/compose.yml` today."
	l := pointer_tester.Checker(
		"pkg/tool/goalertlogd",
		"strata/container",
		"strata/container/alfa",
		"strata/container/alfa/compose.yaml",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf(
				"---\nbase: pkg/tool/goalertlogd, strata/container\n---\n%s\n",
				line,
			),
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadConcerns("Alfa", 4, line),
		"",
		l,
	)
}

func TestPointersUnprefixedSymbolIsPath(t *testing.T) {
	line := "Wraps `pkg/provision/salt.Client` today."
	l := pointer_tester.Checker("pkg/provision/salt")(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersSymbolPlaceholderIgnored(t *testing.T) {
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader("Write `go:pkg/<name>/Symbol` in examples.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersUppercaseFileDead(t *testing.T) {
	line := "Build from `strata/container/gopostgres/Containerfile` here."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersUppercaseFileExisting(t *testing.T) {
	l := pointer_tester.Checker("strata/container/gopostgres/Containerfile")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Build from `strata/container/gopostgres/Containerfile` here.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersSiblingSymbolIgnored(t *testing.T) {
	l := pointer_tester.Checker("../github/soil/pkg/provision/salt")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Wraps `go:../github/soil/pkg/provision/salt.Client` today.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersSiblingSymbolDeadPackage(t *testing.T) {
	line := "Wraps `go:../github/soil/pkg/gone/Thing` today."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersReceiverMethodIgnored(t *testing.T) {
	l := pointer_tester.Checker("pkg/lint/pointer")(
		stringsConstant.UpperAlfa,
		strings.NewReader("Calls `go:pkg/lint/pointer/Names.ResolvePackage`.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersQualifiedMethodIgnored(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision/salt")(
		stringsConstant.UpperAlfa,
		strings.NewReader("Calls `go:pkg/provision/salt.Client.Login`.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersSymbolExpansion(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision/salt", "pkg/provision/ansible")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Both `go:pkg/provision/{salt,ansible}.Client` work.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersLiveSymbolInLivePackage(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision/salt")(
		stringsConstant.UpperAlfa,
		strings.NewReader("Renamed away: `go:pkg/provision/salt.Gone`.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersBasedAnchorResolves(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision", "pkg/provision/salt")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/provision\n---\nConnectors live in `salt/` today.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersBasedInteriorDead(t *testing.T) {
	line := "Connectors live in `salt/gone.go` today."
	l := pointer_tester.Checker("pkg/provision", "pkg/provision/salt")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nbase: pkg/provision\n---\n%s\n", line),
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadConcerns("Alfa", 4, line),
		"",
		l,
	)
}

func TestPointersBasedForeignNamespaceIgnored(t *testing.T) {
	l := pointer_tester.Checker("pkg/provision")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/provision\n---\nGroups sit in `keepass/General` outside.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersBasedSibling(t *testing.T) {
	l := pointer_tester.Checker(
		"../github/soil/pkg/tool/gomaintlogd",
		"../github/soil/pkg/tool/gomaintlogd/server",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nbase: ../github/soil/pkg/tool/gomaintlogd\n---\nHandlers sit in `server/` today.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersBasedSiblingFile(t *testing.T) {
	l := pointer_tester.Checker(
		"../github/soil/pkg/tool/gomaintlogd",
		"../github/soil/pkg/tool/gomaintlogd/face",
		"../github/soil/pkg/tool/gomaintlogd/face/source.go",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nbase: ../github/soil/pkg/tool/gomaintlogd\n---\nThe reader lives in `face/source.go` today.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersBasedSiblingInteriorDead(t *testing.T) {
	line := "Handlers sit in `server/gone.go` today."
	l := pointer_tester.Checker(
		"../github/soil/pkg/tool/gomaintlogd",
		"../github/soil/pkg/tool/gomaintlogd/server",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf(
				"---\nbase: ../github/soil/pkg/tool/gomaintlogd\n---\n%s\n",
				line,
			),
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadConcerns("Alfa", 4, line),
		"",
		l,
	)
}

func TestPointersCommandResolves(t *testing.T) {
	checker, count := pointer_tester.Counting(".claude/skills/sign-firefox")
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader("Invoke `/sign-firefox` before pushing.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersCommandPluginResolves(t *testing.T) {
	checker, count := pointer_tester.Counting(
		"../soil/strata/plugin/soil/skills/lint",
	)
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader("Invoke `/soil:lint` after the scope.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersCommandDead(t *testing.T) {
	line := "Invoke `/ghost-skill` before pushing."
	checker, count := pointer_tester.Counting()
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
	assert.Integer(t, 0, *count)
}

func TestPointersCommandHarness(t *testing.T) {
	checker, count := pointer_tester.Counting()
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader("New tools need a `/mcp` reconnect.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersCommandPluginInRepository(t *testing.T) {
	checker, count := pointer_tester.Counting("strata/plugin/soil/skills/lint")
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader("Invoke `/soil:lint` after the scope.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersRouteUnchecked(t *testing.T) {
	checker, count := pointer_tester.Counting()
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"The service answers on `route:/health` for probes.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Integer(t, 1, *count)
}

func TestPointersRouteWithoutSpecificationDead(t *testing.T) {
	line := "Lists at `route:/api/alerts`."
	l := pointer_tester.Checker("pkg/tool/alfa")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nbase: pkg/tool/alfa\n---\n%s\n", line),
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadConcerns("Alfa", 4, line),
		"",
		l,
	)
}

func TestPointersRouteLive(t *testing.T) {
	l := pointer_tester.Routes(
		map[string][]string{
			"pkg/tool/alfa": {"/api/alerts", "/api/alerts/{name}"},
		},
		"pkg/tool/alfa",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/tool/alfa\n---\nLists at `route:/api/alerts`, one at `route:/api/alerts/{identifier}`, filtered by `route:/api/alerts?name=x` and `route:/api/...`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersRouteDead(t *testing.T) {
	line := "Lists at `route:/api/ghosts` today."
	l := pointer_tester.Routes(
		map[string][]string{"pkg/tool/alfa": {"/api/alerts"}},
		"pkg/tool/alfa",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nbase: pkg/tool/alfa\n---\n%s\n", line),
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadConcerns("Alfa", 4, line),
		"",
		l,
	)
}

func TestPointersRouteLiteralLive(t *testing.T) {
	l := pointer_tester.Literals(
		map[string][]string{
			"pkg/tool/alfa": {
				`WidgetsPath = "/widgets"`,
				`mux.HandleFunc("/memories/", list)`,
			},
		},
		nil,
		nil,
		"pkg/tool/alfa",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/tool/alfa\n---\nPages at `route:/widgets` and `route:/memories/{identifier}`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersRouteLiteralDead(t *testing.T) {
	line := "Pages at `route:/ghosts` today."
	l := pointer_tester.Literals(
		map[string][]string{"pkg/tool/alfa": {`WidgetsPath = "/widgets"`}},
		nil,
		nil,
		"pkg/tool/alfa",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nbase: pkg/tool/alfa\n---\n%s\n", line),
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadConcerns("Alfa", 4, line),
		"",
		l,
	)
}

func TestPointersRouteLiteralInSharedTree(t *testing.T) {
	l := pointer_tester.Literals(
		map[string][]string{
			"pkg/tool/alfa":        {`WidgetsPath = "/widgets"`},
			"../github/shared/pkg": {`HealthPath = "/health"`},
		},
		nil,
		[]string{"../github/shared/pkg"},
		"pkg/tool/alfa",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nbase: pkg/tool/alfa\n---\nProbes `route:/health`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersImageLive(t *testing.T) {
	l := pointer_tester.Literals(
		map[string][]string{"": {"image: ghcr.io/example/alfa:v1"}},
		nil,
		nil,
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader("Runs `ghcr.io/example/alfa`.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersMediaLive(t *testing.T) {
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Serves `media:text/html`, `media:application/vnd.api+json` and `media:haptics/ivs`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersMediaUnregistered(t *testing.T) {
	line := "Serves `media:txt/html`, `media:text/` and `media:text/html/full`."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.UnregisteredMedia("Alfa", line, 3),
		"",
		l,
	)
}

func TestPointersMediaShapedShortHinted(t *testing.T) {
	line := "The stale path served empty `text/html` today."
	l := pointer_tester.Implicit(
		[]string{constant.PackageDirectory},
		constant.PackageDirectory,
		"pkg/text",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadHinted(
			"Alfa",
			1,
			line,
			"Referenced path does not exist - if a media type, write media:text/html",
		),
		"",
		l,
	)
}

func TestPointersMediaShapedShortHintedUnderBase(t *testing.T) {
	line := "The stale path served empty `text/html` today."
	l := pointer_tester.Checker("pkg/tool/alfa", "pkg/tool/alfa/text")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nbase: pkg/tool/alfa\n---\n%s\n", line),
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadHinted(
			"Alfa",
			4,
			line,
			"Referenced path does not exist - if a media type, write media:text/html",
		),
		"",
		l,
	)
}

func TestPointersBareSlashDead(t *testing.T) {
	line := "Lists at `/api/alerts` and `/api/alerts/{name}`."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.BareSlash("Alfa", line, 2),
		"",
		l,
	)
}

func TestPointersSystemUnchecked(t *testing.T) {
	checker, seen := pointer_tester.Recording()
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Mounts `/proc/`, reads `/etc/app.conf`, runs `path:/bin/true`, forgets `/bin/false`, pulls `ghcr.io/example/image`, imports `\"github.com/example/module\"`, rewrites with `s/a/b/`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
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
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("---\n%s\n---\nProse.\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadConcerns("Alfa", 2, line),
		"",
		l,
	)
}

func TestPointersExpansionExisting(t *testing.T) {
	l := pointer_tester.Checker(
		"doc/ai/spec/naming.md",
		"doc/ai/spec/build.md",
	)(
		stringsConstant.UpperAlfa,
		strings.NewReader("Read `doc/ai/spec/{naming,build}.md` first.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersExpansionDead(t *testing.T) {
	line := "Read `doc/ai/spec/{naming,ghost}.md` first."
	l := pointer_tester.Checker("doc/ai/spec/naming.md")(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersSiblingExisting(t *testing.T) {
	l := pointer_tester.Checker("../../github/soil/doc/ai/spec/naming.md")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Shared specs at `../../github/soil/doc/ai/spec/naming.md`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersSiblingDead(t *testing.T) {
	line := "Shared specs at `../github/soil/doc/ai/spec/naming.md`."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 1),
		"",
		l,
	)
}

func TestPointersSiblingGitignored(t *testing.T) {
	l := pointer_tester.Gitignored("../soil/tmp/harness")(
		stringsConstant.UpperAlfa,
		strings.NewReader("Harness at `../soil/tmp/harness/` today.\n"),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersMultiple(t *testing.T) {
	line := "Read `doc/ai/spec/ghost.md` and `pkg/gone`."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.Dead("Alfa", line, 2),
		"",
		l,
	)
}

func TestPointersPluginRoot(t *testing.T) {
	l := pointer_tester.Checker("doc/ai/runbook/lint.md")(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Read `${CLAUDE_PLUGIN_ROOT}/doc/ai/runbook/lint.md`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersBareLinkDead(t *testing.T) {
	line := "See [ghost](ghost.md) for details."
	l := pointer_tester.Checker()(
		"doc/guide/index.md",
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
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
	lint_tester.New(t).Report("doc/guide/index.md", false, nil, "", l)
}

func TestPointersHostDeclared(t *testing.T) {
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nhosts: alfa.example\n---\nStatus at `https://alfa.example/api/v2/status` and [console](http://alfa.example:8080/console).\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersHostDeclaredList(t *testing.T) {
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\nhosts:\n- alfa.example\n- bravo.example\n---\nMirrors `https://alfa.example/` and `https://user:secret@bravo.example/`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersHostUndeclared(t *testing.T) {
	line := "Status at `https://ghost.example/api/v2/status` today."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(fmt.Sprintf("%s\n", line)),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.UndeclaredHost("Alfa", 1, line, 1),
		"",
		l,
	)
}

func TestPointersHostUndeclaredLink(t *testing.T) {
	line := "See [ticket](https://ghost.example/browse/X-1) first."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nhosts: other.example\n---\n%s\n", line),
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.UndeclaredHost("Alfa", 4, line, 1),
		"",
		l,
	)
}

func TestPointersHostImplicit(t *testing.T) {
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Listens on `http://localhost:8090/mcp` and `http://127.0.0.1:9191`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersHostPlaceholder(t *testing.T) {
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Fetch `https://$DD_HOST/api/v2/products/` or `https://<master>:443`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
}

func TestPointersSchemeUnchecked(t *testing.T) {
	checker, seen := pointer_tester.Recording()
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"Resource `guide://summary`, share `smb://files.example/team`, link `tcp://10.0.0.1:8080`.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Strings(
		t,
		[]string{
			"Alfa:1 scheme guide://summary",
			"Alfa:1 scheme smb://files.example/team",
			"Alfa:1 scheme tcp://10.0.0.1:8080",
		},
		*seen,
	)
}

func TestPointersCommandDeclared(t *testing.T) {
	checker, count := pointer_tester.Counting()
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\ncommands:\n- /nodes\n- /migration\n---\nThe ACL applies to `/nodes` and the `/migration` znode stays absent.\n",
		),
	)
	lint_tester.New(t).Report("Alfa", false, nil, "", l)
	assert.Integer(t, 0, *count)
}

func TestPointersCommandDeclaredScalar(t *testing.T) {
	checker, count := pointer_tester.Counting()
	l := checker(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			"---\ncommands: /compact\n---\nThe harness handles `/compact` itself and `/ghost` stays dead.\n",
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.DeadConcerns(
			"Alfa",
			4,
			"The harness handles `/compact` itself and `/ghost` stays dead.",
		),
		"",
		l,
	)
	assert.Integer(t, 0, *count)
}

func TestPointersHostMalformed(t *testing.T) {
	line := "Broken at `https://alfa.example:port/status` today."
	l := pointer_tester.Checker()(
		stringsConstant.UpperAlfa,
		strings.NewReader(
			fmt.Sprintf("---\nhosts: alfa.example\n---\n%s\n", line),
		),
	)
	lint_tester.New(t).Report(
		"Alfa",
		true,
		pointer_tester.UndeclaredHost("Alfa", 4, line, 1),
		"",
		l,
	)
}
