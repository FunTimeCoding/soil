package constant

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

const (
	HostEnvironment    = "MATTERMOST_HOST"
	TokenEnvironment   = "MATTERMOST_TOKEN"
	TeamEnvironment    = "MATTERMOST_TEAM"
	ChannelEnvironment = "MATTERMOST_CHANNEL"

	PerPage    int = 1000
	MaxPerPage int = 200

	// The server caps posts-since responses near a thousand posts,
	// keeping the oldest. A chunk at or above the threshold is
	// treated as capped and fetching continues past its newest
	// post, up to the chunk limit (a bound for pathological
	// windows, about twenty thousand posts).
	SinceChunkThreshold int = 1000
	SinceChunkLimit     int = 20

	EmptyEntityTag = ""

	PostField = "post"
)

// Emoji
const (
	CheckMark    = "white_check_mark"       // done
	Construction = "construction"           // in progress
	Hourglass    = "hourglass_flowing_sand" // waiting, pending resolve
	Repeat       = "repeat"                 // forwarded
	Thread       = "thread"                 // belongs to above thread
)

var (
	ErrorNotConfigured = errors.New("not configured")
	ErrorNotFound      = errors.New("not found")
	Format             = option.ExtendedColor.Copy()
)
