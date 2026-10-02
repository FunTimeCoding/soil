package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"regexp"
)

const (
	ModFile = "go.mod"
	SumFile = "go.sum"

	RuntimeOld    = "runtime_old"
	RuntimeNewer  = "runtime_newer"
	PanicOccurred = "panic_occurred"

	LicenseMissing = "license_missing"
	LicenseUnknown = "license_unknown"
	Permissive     = "permissive"
	WeakCopyleft   = "weak_copyleft"
	Copyleft       = "copyleft"

	AllPackages          = "./..."
	DependenciesArgument = "-deps"
	FormatArgument       = "-f"
	ModuleTemplate       = "{{if not .Standard}}{{.Module.Main}} {{.Module.Path}} {{.Module.Version}} {{.Module.Dir}}{{end}}"
)

const VersionSkipEnvironment = "VERSION_SKIP"

var (
	Format         = constant.ColorFormat.Copy()
	DeadTagPattern = regexp.MustCompile(
		`(\S+)@(\S+): reading .+ unknown revision`,
	)
	LicenseFileNames   = []string{"LICENSE", "LICENCE", "COPYING"}
	PermissiveLicenses = []string{
		"MIT",
		"BSD",
		"0BSD",
		"Apache",
		"ISC",
		"Unlicense",
		"CC0",
		"Zlib",
		"BSL",
	}
	WeakCopyleftLicenses = []string{"MPL", "LGPL", "EPL", "CDDL"}
	CopyleftLicenses     = []string{"GPL", "AGPL"}
)
