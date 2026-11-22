package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"regexp"
	"time"
)

var Identity = identity.New(
	"gosecurityd",
	"Security advisory bridge",
	"gosecurityd",
).WithInstructions(
	"Security advisories and their remediation. Systems report their installed packages; gosecurityd matches them against Debian's and Freexian's trackers and scores each vulnerability from the National Vulnerability Database. Read gosecurityd://guide/matching before judging a system (coverage, the clean rule, fix kinds, the kernel) and gosecurityd://guide/remediation before working through the remediation or recording a decision. Start with list_remediation for what to do; list_systems, list_findings and list_upgrades for the fleet; get_vulnerability, get_debian_status and get_affected_systems for one vulnerability.",
)

const (
	GetVulnerability  = "get_vulnerability"
	GetDebianStatus   = "get_debian_status"
	FlagVulnerability = "flag_vulnerability"
	ListFlags         = "list_flags"
	MitigateFlag      = "mitigate_flag"
	ListSystems       = "list_systems"
	GetSystemFindings = "get_system_findings"
	ListFindings      = "list_findings"
	ListRemediation   = "list_remediation"
)

const (
	NationalDatabasePeer = "national vulnerability database upstream"
	SecurityTrackerPeer  = "debian security tracker upstream"
	ExtendedTrackerPeer  = "freexian extended tracker upstream"
	AnsiblePeer          = "goansibled"
)

const (
	ProbeInterval    = 24 * time.Hour
	ProbeHostCommand = "security_probe"
	ProbeSource      = "gosecurityd"
	ProbeSkipped     = "probe %s skipped: %s: %s"
	ProbeFailed      = "probe %s: %w"
)

const (
	DebianCache   = "debian-tracker.json"
	ExtendedCache = "freexian-tracker.json"
)

var (
	SupportedReleases = []string{"bookworm", "trixie", "forky", "sid"}
	ExtendedReleases  = []string{"stretch", "buster", "bullseye"}
)

const RefreshInterval = time.Hour
const (
	ScoreInterval = time.Hour
	ScoreOverlap  = time.Hour
)
const (
	StaleAfter        = 72 * time.Hour
	TrackerStaleAfter = 24 * time.Hour
)

const (
	StaleArgument        = "stale-after"
	StaleUsage           = "Snapshot age after which a system no longer reads clean"
	TrackerStaleArgument = "tracker-stale-after"
	TrackerStaleUsage    = "Tracker age after which no system on it reads clean"
)
const FlagKind = "flag"
const (
	FlagSubjectColumn = "subject"
	FlagEarlierColumn = "vulnerability"
	FlagEarlierIndex  = "idx_security_flags_vulnerability"
)
const (
	RemediationVulnerability = "vulnerability"
	RemediationUpgrade       = "upgrade"
	RemediationRelease       = "release"
	RemediationReboot        = "reboot"
	RemediationKernel        = "kernel"
	SubjectSeparator         = ":"
	Unimportant              = "unimportant"
	RemediationTitle         = "Remediation"
	RemediationPath          = "/remediation"
)
const (
	FlagsTitle   = "Flags"
	FlagsPath    = "/"
	LookupTitle  = "Lookup"
	LookupPath   = "/lookup"
	FlagPath     = "/flag"
	MitigatePath = "/mitigate"
	NoteField    = "note"
	FlagFormMark = "flag-form"
	FlagsMark    = "flags-table"
	SystemsMark  = "affected-systems"
	FlagOpenPath = "/flag/open"
	FlagShutPath = "/flag/close"
	RowPrefix    = "flag-"
	ShowField    = "show"
	FlagColumns  = 4
)

const (
	FixDebian = "debian fix"
	FixLocal  = "local action"
	FixAbsent = "no tracker data"
)

const (
	CriticalScore = 9.0
	HighScore     = 7.0
)
const SystemKind = "system"
const (
	CoverageSupported = "supported"
	CoverageTracked   = "unsupported, tracked"
	CoverageUntracked = "unsupported, untracked"
)

const (
	ShowSupported = "supported"
	ShowTracked   = "tracked"
	ShowUntracked = "untracked"
)

var CoverageFilters = map[string]string{
	"":            "",
	ShowSupported: CoverageSupported,
	ShowTracked:   CoverageTracked,
	ShowUntracked: CoverageUntracked,
}

const (
	ListUpgrades       = "list_upgrades"
	GetUpgrade         = "get_upgrade"
	GetAffectedSystems = "get_affected_systems"
	GetScoreStatus     = "get_score_status"
)

const (
	KindUpgrade = "upgrade fixes it"
	KindPending = "debian fix not in candidate"
	KindPaid    = FindingPaid
	KindOpen    = "no fix yet"
)

const (
	KindReboot        = "reboot fixes it"
	KindKernelUpgrade = "upgrade and reboot fix it"
)

var KindOrder = []string{
	KindReboot,
	KindUpgrade,
	KindKernelUpgrade,
	KindPending,
	KindPaid,
	KindOpen,
}

const (
	KernelCurrent = "up to date"
	KernelReboot  = KindReboot
	KernelUpgrade = "upgrade and reboot"
	KernelStuck   = "no upgrade path"
	KernelPaid    = FindingPaid
	KernelUnknown = "no kernel data"
)

var KernelStateOrder = []string{
	KernelReboot,
	KernelUpgrade,
	KernelStuck,
	KernelPaid,
	KernelCurrent,
	KernelUnknown,
}

const (
	ListsOldAfter            = 30 * 24 * time.Hour
	ReasonNoBackports        = "no backports source"
	ReasonListsOld           = "package lists %d days old"
	ReasonOffer              = "sources offer up to %s"
	ReasonUnknown            = "sources not reported yet"
	ReasonNoMetapackage      = "kernel metapackage not installed"
	ReasonNoMetapackageOffer = "kernel metapackage not installed (sources offer up to %s)"
	BackportsOrigin          = "backports"
)

const (
	KernelImagePrefix = "linux-image-"
	BackportsLabel    = " (backports)"
)

var BackportsPattern = regexp.MustCompile(`~bpo(\d+)\+\d+$`)
var ReleaseNumbers = map[string]int{
	"stretch":  9,
	"buster":   10,
	"bullseye": 11,
	"bookworm": 12,
	"trixie":   13,
	"forky":    14,
}

const (
	UpgradesTitle    = "Upgrades"
	UpgradesPath     = "/upgrades"
	UpgradeOpenPath  = "/upgrade/open"
	UpgradeShutPath  = "/upgrade/close"
	UpgradeRowPrefix = "upgrade-"
	UpgradeColumns   = 7
	UpgradeKind      = "upgrade"
	KernelFilter     = "kernel"
	KernelSource     = "linux"
	KernelPrefix     = "linux-"
	AptUpgrade       = "apt-get install --only-upgrade"
	ReleaseParameter = "release"
	SourceParameter  = "source"
)

const (
	SystemsTitle = "Systems"
	SystemsPath  = "/systems"
	SystemPath   = "/system"
)

const (
	LabelField     = "label"
	SearchField    = "search"
	LabelSeparator = ":"
	SearchLabel    = "Host name"
)

const (
	VulnerabilitiesTitle  = "Vulnerabilities"
	VulnerabilitiesPath   = "/vulnerabilities"
	VulnerabilitiesHeader = "Vulnerabilities"
	ScoreField            = "score"
	HighestScore          = 10.0
)

var (
	CriticalChips = [][2]string{{"All scores", ""}, {"9.0 or higher", "9"}}
	FindingChips  = [][2]string{
		{"9.0 or higher", "9"},
		{"7.0 or higher", "7"},
		{"All scores", "0"},
	}
)

const ScoreDescription = "The lowest score to list, e.g. 9 for critical; empty or 0 lists everything"
const (
	FindingDebian = "debian"
	FindingPaid   = "paid support only"
	FindingNone   = "no fix"
)

var (
	SystemNamePattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]*$`)
	RowUnsafePattern  = regexp.MustCompile(`[^a-z0-9-]`)
)
