package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goffrelay",
	"Firefox native messaging relay to the gofirefoxd bridge",
	"goffrelay [flags] [manifest] [extension]",
)

const (
	Extension     = "{6a4363a0-49d6-415b-9ea6-e0feda0b78f5}"
	ExtensionFlag = "extension"

	HeaderLength = 4
	FrameLimit   = 64 * 1024 * 1024

	Scheme = "ws://"
)
