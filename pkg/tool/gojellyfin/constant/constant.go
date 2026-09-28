package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gojellyfin",
	"Jellyfin media server - libraries, items, playback sessions",
	"gojellyfin <command>",
)

const (
	HostEnvironment     = "GOJELLYFIN_HOST"
	PortEnvironment     = "GOJELLYFIN_PORT"
	InsecureEnvironment = "GOJELLYFIN_INSECURE"
	TokenEnvironment    = "GOJELLYFIN_TOKEN"
)
