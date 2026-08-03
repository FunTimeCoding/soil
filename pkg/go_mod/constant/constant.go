package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"regexp"
)

const ModFile = "go.mod"
const (
	RuntimeOld    = "runtime_old"
	RuntimeNewer  = "runtime_newer"
	PanicOccurred = "panic_occurred"
)

const VersionSkipEnvironment = "VERSION_SKIP"

var Format = constant.ColorFormat.Copy()
var DeadTagPattern = regexp.MustCompile(
	`(\S+)@(\S+): reading .+ unknown revision`,
)
