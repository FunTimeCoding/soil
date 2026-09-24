package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gohook",
	"Git hook manager with path-conditional jobs",
	"gohook <command> [flags]",
)

const (
	RootFile  = ".gohook.yaml"
	ToolFile  = "strata/tool/gohook.yaml"
	Extension = ".yaml"

	SkipEnvironment = "GOHOOK"
	SkipValue       = "0"

	Stub = `#!/bin/sh
# gohook
[ "$GOHOOK" = "0" ] && exit 0
exec gohook run %s -- "$@"
`
	StubMarker = "# gohook"
	StubMode   = 0o755

	Install   = "install"
	Uninstall = "uninstall"
	Run       = "run"
	RunUsage  = "run <hook> [-- hook arguments]"
	List      = "list"

	MissingConfiguration = "no configuration: expected %s or %s under %s"
	HooksPathRedirected  = "core.hooksPath is %q so git would never run stubs in %s; unset it (git config --unset core.hooksPath, check --global too) and install again"
	UnknownHook          = "unknown hook %q in %s, expected one of %s"
	MissingRun           = "job %d under %s has no run"
	JobFailed            = "%s: job %d failed: %s\n"
	JobModified          = "%s: job %d modified files that are not staged, review and stage them: %s\n%s\n"
	NoJobs               = "%s: no jobs\n"
	NothingPushed        = "%s: nothing pushed, skipping\n"
	Skipped              = "%s: skip %s\n"
	Running              = "%s: run %s\n"
	ListHook             = "%s (%s)\n"
	ListJob              = "  %s\n"
	ListJobPaths         = "  %s  [%s]\n"
	Installed            = "installed %s\n"
	Removed              = "removed %s\n"
	Kept                 = "kept %s (not written by gohook)\n"
)
