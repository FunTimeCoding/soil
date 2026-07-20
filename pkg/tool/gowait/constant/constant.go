package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"gowait",
	"Service readiness waiter",
	"gowait [flags]",
)

const Interval = 10 * time.Second
