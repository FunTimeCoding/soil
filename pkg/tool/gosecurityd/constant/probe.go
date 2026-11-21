package constant

import "regexp"

const (
	ReleaseSection    = "##release"
	KernelSection     = "##kernel"
	PackagesSection   = "##packages"
	UpgradableSection = "##upgradable"
	ListsSection      = "##lists"
	PolicySection     = "##kernel-policy"
	SectionPrefix     = "##"
)

const ProbeCommand = `echo '##release'; cat /etc/os-release 2>/dev/null || cat /etc/redhat-release 2>/dev/null; echo '##kernel'; uname -r; echo '##packages'; dpkg-query -W -f '${db:Status-Abbrev}\t${binary:Package}\t${Version}\t${source:Package}\t${source:Version}\n' 2>/dev/null; echo '##upgradable'; apt list --upgradable 2>/dev/null; echo '##lists'; find /var/lib/apt/lists -maxdepth 1 -name '*Release' -printf '%T@\n' 2>/dev/null | sort -n | tail -1; echo '##kernel-policy'; apt-cache policy linux-image-$(dpkg --print-architecture 2>/dev/null) 2>/dev/null`

var (
	CodenamePattern  = regexp.MustCompile(`\(([a-z]+)\)`)
	RedHatPattern    = regexp.MustCompile(`^(\w+)(?: Linux)? release (\d+)`)
	OfferPattern     = regexp.MustCompile(`^\s+(?:\*\*\*\s+)?(\S+)\s+(-?\d+)$`)
	OriginPattern    = regexp.MustCompile(`^\s+(-?\d+)\s+(\S.*)$`)
	InstalledPattern = regexp.MustCompile(`^\s+Installed:\s+(\S+)`)
)

const NotInstalled = "(none)"
const (
	ReleaseIdentifier = "ID"
	ReleaseCodename   = "VERSION_CODENAME"
	ReleaseVersion    = "VERSION"
	ReleaseNumber     = "VERSION_ID"
	DebianIdentifier  = "debian"
	InstalledStatus   = 'i'
	PackageFields     = 5
	IndexSuffix       = "Packages"
	SuiteSeparator    = "/"
	VersionSeparator  = "."
)
