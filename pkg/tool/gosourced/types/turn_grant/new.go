package turn_grant

import "time"

func New(
	ahead int,
	waited time.Duration,
	release func(),
) *Grant {
	return &Grant{Ahead: ahead, Waited: waited, Release: release}
}
