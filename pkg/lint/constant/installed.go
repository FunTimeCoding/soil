package constant

const (
	StaleBinaryKey  = "stale_binary"
	StaleBinaryText = "Installed binary predates changes to its source - in %s run gobuild --copy-to-bin %s"
	DirtyBinaryKey  = "dirty_binary"
	DirtyBinaryText = "Installed binary was built from uncommitted changes - in %s run gobuild --copy-to-bin %s"

	LinkerFlagsSetting   = "-ldflags"
	MainVersionVariable  = "main.Version"
	MainGitHashVariable  = "main.GitHash"
	CommandDirectory     = "cmd"
	CommandLineArguments = "command-line-arguments"
	PackageListTemplate  = "{{.ImportPath}}\t{{.Dir}}\t{{join .Deps \" \"}}"
	FormatArgument       = "-f"
	ModuleListTemplate   = "{{.Path}} {{if .Replace}}{{.Replace.Version}}{{else}}{{.Version}}{{end}}"
	AllModules           = "all"
	Head                 = "HEAD"
	CommitSuffix         = "^{commit}"
)
