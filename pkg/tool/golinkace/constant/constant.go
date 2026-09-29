package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"golinkace",
	"LinkAce bookmark management CLI",
	"golinkace [command]",
)

const (
	DaemonHostEnvironment     = "GOLINKACE_HOST"
	DaemonPortEnvironment     = "GOLINKACE_PORT"
	DaemonInsecureEnvironment = "GOLINKACE_INSECURE"
	DaemonTokenEnvironment    = "GOLINKACE_TOKEN"
	LinkAceHostEnvironment    = "LINKACE_HOST"

	ListFlag = "list"
)
