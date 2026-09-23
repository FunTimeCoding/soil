package constant

import "regexp"

const (
	ConventionWordKey  = "convention_word"
	ConventionWordText = "Bare convention directory - anchor it with base:, expand it, or write <path>/<name>/"

	BareSlashKey  = "bare_slash"
	BareSlashText = "Bare /api/ span - write route:/api/... so it checks against the openapi specification"

	PluginRootPrefix = "${CLAUDE_PLUGIN_ROOT}/"

	SchemeGo    = "go:"
	SchemePath  = "path:"
	SchemeRoute = "route:"

	BaseKey = "base"

	ConfigurationEnvironment = "GOLINT_CONFIGURATION"

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
		VerdictDead:       DeadPointerKey,
		VerdictAbsolute:   AbsolutePointerKey,
		VerdictConvention: ConventionWordKey,
		VerdictBareSlash:  BareSlashKey,
	}

	VerdictTexts = map[Verdict]string{
		VerdictDead:       DeadPointerText,
		VerdictAbsolute:   AbsolutePointerText,
		VerdictConvention: ConventionWordText,
		VerdictBareSlash:  BareSlashText,
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
