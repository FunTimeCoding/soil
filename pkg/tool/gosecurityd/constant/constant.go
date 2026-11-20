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
	"Security advisories. get_vulnerability answers a CVE's description, score and references from the National Vulnerability Database. get_debian_status answers, per Debian source package and release, whether the CVE is fixed and in which package version - Debian backports fixes without changing the upstream version, so use it before judging a Debian host by version ranges. flag_vulnerability records a CVE as a concern to mitigate, list_flags shows the open concerns, mitigate_flag closes one with how it was mitigated.",
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
)

const (
	NationalDatabasePeer = "nvd upstream"
	SecurityTrackerPeer  = "debian security tracker upstream"
	ExtendedTrackerPeer  = "freexian extended tracker upstream"
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
	FlagsTitle   = "Flags"
	FlagsPath    = "/"
	LookupTitle  = "Lookup"
	LookupPath   = "/lookup"
	FlagPath     = "/flag"
	MitigatePath = "/mitigate"
	NoteField    = "note"
	FlagFormMark = "flag-form"
	FlagsMark    = "flags-table"
	FlagOpenPath = "/flag/open"
	FlagShutPath = "/flag/close"
	RowPrefix    = "flag-"
	ShowField    = "show"
	FlagColumns  = 5
)

const (
	FixDebian = "debian fix"
	FixOwn    = "own work"
	FixAbsent = "no tracker data"
)

const CriticalScore = 9.0
const SystemKind = "system"
const (
	CoverageSupported = "supported"
	CoverageTracked   = "unsupported, tracked"
	CoverageUntracked = "unsupported, untracked"
)

var CoverageFilters = map[string]string{
	"":          "",
	"supported": CoverageSupported,
	"tracked":   CoverageTracked,
	"untracked": CoverageUntracked,
}

const (
	SystemsTitle = "Systems"
	SystemsPath  = "/systems"
	SystemPath   = "/system"
)

const (
	FindingDebian = "debian"
	FindingPaid   = "paid support only"
	FindingNone   = "no fix"
)

var SystemNamePattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]*$`)
