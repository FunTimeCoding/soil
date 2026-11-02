package constant

import "github.com/funtimecoding/soil/pkg/identity"

const (
	HostEnvironment     = "GORAIDPARSE_HOST"
	PortEnvironment     = "GORAIDPARSE_PORT"
	InsecureEnvironment = "GORAIDPARSE_INSECURE"
	TokenEnvironment    = "GORAIDPARSE_TOKEN"
)

const (
	NoFiles    = "no files to parse"
	ReportName = "ONYX Log %s.html"
	ReportTime = "2006-01-02 15-04"
)

var Identity = identity.New(
	"goraidparsed",
	"RAID log parser service",
	"goraidparsed [flags]",
)
