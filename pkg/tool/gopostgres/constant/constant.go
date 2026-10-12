package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gopostgres",
	"PostgreSQL query CLI for gopostgresd",
	"gopostgres [command]",
)

const (
	HostEnvironment     = "GOPOSTGRES_HOST"
	PortEnvironment     = "GOPOSTGRES_PORT"
	InsecureEnvironment = "GOPOSTGRES_INSECURE"
	TokenEnvironment    = "GOPOSTGRES_TOKEN"
)
