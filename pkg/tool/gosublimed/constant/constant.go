package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gosublimed",
	"Sublime Text editor bridge",
	"gosublimed",
).WithInstructions(
	"Sublime Text editor - read, edit, and create buffers. Scratch buffers persist across plugin reloads and Sublime restarts.",
)

const (
	ListViews  = "list_views"
	ReadView   = "read_view"
	EditView   = "edit_view"
	CreateView = "create_view"
	SaveView   = "save_view"
	CloseView  = "close_view"
	OpenFile   = "open_file"
)
