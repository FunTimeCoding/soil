package constant

import consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"

const (
	HostEnvironment  = "LINKACE_HOST"
	TokenEnvironment = "LINKACE_TOKEN"

	DefaultPerPage = 24
	BasePath       = "api/v2"

	ListSubject = "list"
	TagSubject  = "tag"
)

var (
	Format = consoleConstant.ExtendedColorFormat.Copy()
)
