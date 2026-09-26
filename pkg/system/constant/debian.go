package constant

import "github.com/funtimecoding/soil/pkg/system/debian/release"

const (
	DebianPackageConfigurationDirectory = "DEBIAN"

	PreseedConfiguration = "preseed.cfg"

	DebianControlFile = "control"

	DebianPackageExtension = ".deb"

	DebianWeb   = "www.debian.org"
	DebianImage = "cdimage.debian.org"
)

var Bookworm = release.New("bookworm", 12, 1)

const ChecksumFile = "SHA256SUMS"
const (
	Systemctl     = "systemctl"
	SystemctlList = "list-units"
	SystemctlShow = "show"

	SystemctlServiceType = "--type=service"
	SystemctlAll         = "--all"
	SystemctlNoLegend    = "--no-legend"
	SystemctlNoPager     = "--no-pager"
	SystemctlPlain       = "--plain"
	SystemctlProperty    = "-p"
	SystemctlValue       = "--value"
	SystemctlFragment    = "FragmentPath"

	DpkgQuery      = "dpkg-query"
	DpkgSearch     = "--search"
	DpkgList       = "--list"
	AptCache       = "apt-cache"
	AptCachePolicy = "policy"
	AptMark        = "apt-mark"
	AptMarkManual  = "showmanual"

	PolicyInstalled  = "Installed:"
	PolicyCurrent    = "***"
	PolicyStatusPath = "/var/lib/dpkg/status"
	PolicyNone       = "(none)"

	OriginMirrorPrefix = "mirror+file:"

	UnitSuffix      = ".service"
	UnitStateActive = "active"
	UnitStateFailed = "failed"
	UnitLoaded      = "loaded"
	WantsSuffix     = ".wants"
	RequiresSuffix  = ".requires"

	MergedPrefix   = "/usr/lib/"
	UnmergedPrefix = "/lib/"
	UsrDirectory   = "/usr"
)

var SystemdDirectories = []string{
	"/usr/lib/systemd/system",
	"/lib/systemd/system",
	"/etc/systemd/system",
	"/run/systemd/generator",
}
