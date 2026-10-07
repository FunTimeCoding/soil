package constant

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"regexp"
)

const (
	ConventionWordKey  = "convention_word"
	ConventionWordText = "Bare convention directory - anchor it with base:, expand it, or write <path>/<name>/"

	BareSlashKey  = "bare_slash"
	BareSlashText = "Bare /api/ span - write route:/api/... so it checks against the openapi specification"

	UndeclaredHostKey  = "undeclared_host"
	UndeclaredHostText = "Locator host not declared in front matter hosts"

	DeadHeadingKey  = "dead_heading"
	DeadHeadingText = "Referenced heading does not exist"

	FragmentTargetKey  = "fragment_target"
	FragmentTargetText = "Fragment on a non-markdown target - only markdown headings can be referenced"

	FragmentSeparator = "#"
	LocatorSeparator  = "://"
	HintSeparator     = " - "
	NearestPrefix     = "nearest: "
	NearestFormat     = "#%s (%q)"
	NearestLimit      = 3

	PluginRootPrefix = "${CLAUDE_PLUGIN_ROOT}/"

	SchemeGo    = "go:"
	SchemePath  = "path:"
	SchemeRoute = "route:"

	BaseKey = "base"

	ConfigurationPath = "strata/tool/golint.yaml"

	PackageDirectory      = "pkg"
	RestSpecificationPath = "generated/server/openapi.yaml"
	RestRoutePrefix       = "/api/"
	RouteEllipsis         = "..."
	SubstitutionPrefix    = "s/"
	CommentPrefix         = "//"

	BraceCharacters = "{}"

	UserPathPrefix = "/Users/"
	HomePathPrefix = "/home/"
	HomePrefix     = "~"
	Quote          = "\""
)

var (
	VerdictKeys = map[Verdict]string{
		VerdictDead:           DeadPointerKey,
		VerdictAbsolute:       AbsolutePointerKey,
		VerdictConvention:     ConventionWordKey,
		VerdictBareSlash:      BareSlashKey,
		VerdictUndeclaredHost: UndeclaredHostKey,
		VerdictDeadHeading:    DeadHeadingKey,
		VerdictFragmentTarget: FragmentTargetKey,
	}

	VerdictTexts = map[Verdict]string{
		VerdictDead:           DeadPointerText,
		VerdictAbsolute:       AbsolutePointerText,
		VerdictConvention:     ConventionWordText,
		VerdictBareSlash:      BareSlashText,
		VerdictUndeclaredHost: UndeclaredHostText,
		VerdictDeadHeading:    DeadHeadingText,
		VerdictFragmentTarget: FragmentTargetText,
	}

	ImplicitHosts = []string{constant.Localhost, constant.Loopback}

	BareNameExtensions = []string{
		library.MarkdownExtension,
		library.GoExtension,
		library.MarkupExtension,
		library.ShortMarkupExtension,
		library.HypertextExtension,
		".py",
		".sh",
		".json",
		".toml",
		".sql",
		".txt",
		".mod",
	}

	MajorSuffix = regexp.MustCompile(`^v[0-9]+$`)
	LineSuffix  = regexp.MustCompile(`:[0-9]+$`)
	Address     = regexp.MustCompile(
		`^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+(:[0-9]*)?$`,
	)

	ConventionDirectories = []string{
		"argument",
		"base",
		"client",
		"constant",
		"convert",
		"face",
		"generated",
		"helper",
		"integration",
		"mock_client",
		"mock_notifier",
		"model",
		"model_context",
		"option",
		"request",
		"response",
		"result",
		"server",
		"service",
		"store",
		"testdata",
		"tool",
		"toolset",
		"types",
		"unit",
		"util",
		"web",
		"web_interface",
		"web_service",
		"worker",
	}

	HarnessCommands = []string{
		"clear",
		"compact",
		"config",
		"help",
		"hooks",
		"mcp",
		"memory",
		"model",
		"reload-plugins",
		"resume",
		"status",
	}

	ImageRegistries = []string{
		"docker.io",
		"gcr.io",
		"ghcr.io",
		"quay.io",
		"registry.gitlab.com",
	}
)
