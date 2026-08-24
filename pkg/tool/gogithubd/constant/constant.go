package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

const PollInterval = 5 * time.Minute

var Identity = identity.New(
	"gogithubd",
	"GitHub container registry exporter",
	"gogithubd [flags]",
)

const (
	OwnerEnvironment = "GITHUB_EXPORTER_OWNER"
)
