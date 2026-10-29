package constant

import "github.com/funtimecoding/soil/pkg/identity"

const (
	HostEnvironment     = "GORAIDPARSE_HOST"
	PortEnvironment     = "GORAIDPARSE_PORT"
	InsecureEnvironment = "GORAIDPARSE_INSECURE"
	TokenEnvironment    = "GORAIDPARSE_TOKEN"
)

var Identity = identity.New(
	"goraidparsed",
	"RAID log parser service",
	"goraidparsed [flags]",
)
