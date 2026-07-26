package constant

const (
	MattermostHostEnvironment    = "MATTERMOST_HOST"
	MattermostTokenEnvironment   = "MATTERMOST_TOKEN"
	MattermostTeamEnvironment    = "MATTERMOST_TEAM"
	MattermostChannelEnvironment = "MATTERMOST_CHANNEL"

	MattermostPerPage    int = 1000
	MattermostMaxPerPage int = 200

	// The server caps posts-since responses near a thousand posts,
	// keeping the oldest. A chunk at or above the threshold is
	// treated as capped and fetching continues past its newest
	// post, up to the chunk limit (a bound for pathological
	// windows, about twenty thousand posts).
	MattermostSinceChunkThreshold int = 1000
	MattermostSinceChunkLimit     int = 20

	MattermostEmptyEntityTag = ""

	MattermostPostField = "post"
	// Emoji
	MattermostCheckMark    = "white_check_mark"       // done
	MattermostConstruction = "construction"           // in progress
	MattermostHourglass    = "hourglass_flowing_sand" // waiting, pending resolve
	MattermostRepeat       = "repeat"                 // forwarded
	MattermostThread       = "thread"                 // belongs to above thread
)
