package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"gochromed",
	"Chromium browser bridge via CDP",
	"gochromed",
).WithInstructions(
	"Chromium/Brave browser - list tabs, read accessibility snapshots, take screenshots, navigate. Attaches to a running browser via Chrome DevTools Protocol.",
)

const (
	TargetTimeout = 10 * time.Second

	DownloadDirectoryEnvironment = "CHROME_DOWNLOAD_DIRECTORY"

	ListTabs   = "list_tabs"
	Snapshot   = "snapshot"
	Screenshot = "screenshot"
	Navigate   = "navigate"
	Evaluate   = "evaluate"
	CloseTab   = "close_tab"
	CreateTab  = "create_tab"
	ReadBody   = "read_body"
	Click      = "click"
	Fill       = "fill"
	Direct     = "direct"
	Wake       = "wake"
	TabHistory = "history"
)
