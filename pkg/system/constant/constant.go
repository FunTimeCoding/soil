package constant

import "time"

const (
	Tilde = "~"

	Linux   = "linux"
	Darwin  = "darwin"
	Windows = "windows"

	AMD64 = "amd64"
	ARM64 = "arm64"

	LinuxAMD64  = "linux-amd64"
	DarwinARM64 = "darwin-arm64"
	DarwinAMD64 = "darwin-amd64"

	CommandPath                 = "cmd"
	BoardPath                   = "board"
	ConfigurationPath           = ".config"
	DocumentPath                = "doc"
	DownloadsPath               = "Downloads"
	FixturePath                 = "fixture"
	HypertextPath               = "hypertext"
	IdeaPath                    = ".idea"
	KubernetesConfigurationPath = ".kube/config"
	MarkdownPath                = "markdown"
	MemoryPath                  = "memory"
	NotationPath                = "notation"
	SearchPath                  = "search"
	StoragePath                 = ".local/share"
	WikiPath                    = "wiki"
	QuerySocketPath             = ".osquery/shell.em"

	Transmission = "tcp"

	UnknownHost = "unknown"

	Retry = 50 * time.Millisecond
)

// File hierarchy system
const (
	Binary    = "bin"
	Library   = "lib"
	Local     = "local"
	Resources = "usr"
	Source    = "src"
	Temporary = "tmp"
)

// macOS specific paths
const (
	MacOSLibrary = "Library"
	MacOsLogs    = "Logs"
)

var (
	OperatingSystems    = []string{Linux, Darwin}
	Architectures       = []string{AMD64, ARM64}
	SystemArchitectures = []string{LinuxAMD64, DarwinARM64, DarwinAMD64}
)
