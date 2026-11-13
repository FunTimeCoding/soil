package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gobreak",
	"Run a file of breaks through go test -overlay and read each run",
	"gobreak [flags] <break file>",
)

const (
	BreakFile      = 0
	BreakFileUsage = "break file, YAML"
)

const (
	ExpectRed     = "red"
	ExpectSurvive = "survive"
)

const (
	ResultRed        = "red"
	ResultPanic      = "panic"
	ResultUnreached  = "unreached"
	ResultSurvived   = "survived"
	ResultUnmatched   = "unmatched"
	ResultBuildFailed = "build failed"
	ResultSetupFailed = "setup failed"
)

const (
	GoCommand    = "go"
	TestCommand  = "test"
	NotationFlag = "-json"
	CountFlag    = "-count=1"
	TagsFlag     = "-tags"
	RunFlag      = "-run"
	OverlayFlag  = "-overlay"
	CoverFlag    = "-coverprofile"
	CoverageFlag = "-coverpkg"
)

const (
	OverlayFile     = "overlay.json"
	CoverFile       = "cover.out"
	TemporaryPrefix = "gobreak-"
	PackagePrefix   = "./"
	PackageHeader   = "# "
	NameSeparator   = "_"
	FileMode        = 0o600
)

const (
	NoRepository     = "no repository found above %s"
	ProblemFormat    = "%s\n"
	BaselineFailed   = "baseline: tests already fail: %s"
	BaselineBuild    = "baseline: build failed: %s"
	BaselineSetup    = "baseline: packages fail without a failing test: %s"
	BaselineNoTests  = "baseline: the tests pattern matches no test"
	BaselinePassed   = "baseline  ok  %d tests\n"
	UnmatchedFormat  = "%s: %s"
	ReachedNever     = "%s:%d never ran"
	UnexpectedMarker = "  (expected %s)"
	LineFormat       = "%s  %s  %s%s"
	LinePrint        = "%s\n"
	Summary          = "%d breaks, %d as expected\n"
	MissedFormat     = "%d breaks missed their expectation\n"
)
