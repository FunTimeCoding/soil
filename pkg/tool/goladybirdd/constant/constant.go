package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"goladybirdd",
	"Ladybird browser bridge via WebDriver BiDi",
	"goladybirdd",
).WithInstructions(
	"Ladybird browser owned by the daemon - list its tabs, navigate, evaluate scripts, take screenshots. Its own profile, apart from the user's browser.",
)

const (
	LadybirdPeer = "ladybird browser"
	Profile      = "goladybirdd"

	DownloadDirectoryEnvironment = "LADYBIRD_DOWNLOAD_DIRECTORY"

	VisibleArgument = "visible"

	ListTabs   = "list_tabs"
	Navigate   = "navigate"
	Evaluate   = "evaluate"
	Screenshot = "screenshot"
	Snapshot   = "snapshot"
	Click      = "click"
	Fill       = "fill"
	Direct     = "direct"

	DocumentRole  = "document"
	SnapshotLimit = 20000

	NavigationWindow = 150 * time.Millisecond
	NavigationLimit  = 30 * time.Second
)
