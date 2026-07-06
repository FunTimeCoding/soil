package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gofirefoxd",
	"Firefox browser tab management bridge",
	"gofirefoxd",
).WithInstructions(
	"Firefox browser - list, read, create, close, and navigate tabs. Tab groups for organization.",
)

const (
	ListTabs     = "list_tabs"
	ReadTab      = "read_tab"
	CreateTab    = "create_tab"
	CloseTab     = "close_tab"
	Navigate     = "navigate"
	NavigateBack = "navigate_back"
	GroupTabs    = "group_tabs"
	UngroupTab   = "ungroup_tab"
	UpdateGroup  = "update_group"

	BridgePortFlag = "bridge-port"
)
