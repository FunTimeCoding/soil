package constant

import "regexp"

const (
	TopFile        = "top.sls"
	InitName       = "init"
	StateExtension = ".sls"
	StateSeparator = "."
	PrivatePrefix  = "_"
)

const (
	StatusReached   = "reached"
	StatusPossible  = "possible"
	StatusUnreached = "unreached"
)

var (
	DefaultEnvironments = []string{"base", "prod"}

	EnvironmentHeader = regexp.MustCompile(`^([A-Za-z0-9_-]+):\s*(#.*)?$`)
	IncludeHeader     = regexp.MustCompile(`^include:\s*(#.*)?$`)
	RequisiteState    = regexp.MustCompile(`-\s*sls:\s*['"]?([^'"\s#]+)`)
	TemplateReference = regexp.MustCompile(
		`\{%-?\s*(?:include|import(?:_yaml|_json|_text)?|from)\s+['"]([^'"]+)['"]`,
	)
	FileReference = regexp.MustCompile(`salt:///?([^\s'"?]+)`)
	Expression    = regexp.MustCompile(`\{\{.*?\}\}`)
	TemplateLine  = regexp.MustCompile(`^\s*\{[%#].*[%#]\}\s*$`)
)
