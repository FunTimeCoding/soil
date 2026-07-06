package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goitermd",
	"iTerm2 terminal workspace bridge",
	"goitermd",
).WithInstructions(
	"iTerm2 terminal workspace - read screens, send text and keys, manage tabs. Check job_name and command_line before sending to avoid unintended commands on production sessions.",
)

const (
	ListSessions = "list_sessions"
	ReadScreen   = "read_screen"
	ReadHistory  = "read_history"
	SendText     = "send_text"
	SendKey      = "send_key"
	SetTabTitle  = "set_tab_title"
	SetTabColor  = "set_tab_color"
	CreateTab    = "create_tab"
)
