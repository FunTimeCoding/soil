package constant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"time"
)

const (
	DeleteLoopArgument = "delete-loop"

	DiscordTokenEnvironment = "DISCORD_TOKEN"

	DiscordPingCommand    = "!ping"
	DiscordCleanCommand   = "!clean"
	DiscordDetailsCommand = "!details"

	DiscordMessageLimit int = 100

	MattermostHostEnvironment     = "MATTERMOST_HOST"
	MattermostTokenEnvironment    = "MATTERMOST_TOKEN"
	MattermostTeamEnvironment     = "MATTERMOST_TEAM"
	MattermostChannelEnvironment  = "MATTERMOST_CHANNEL"
	MattermostInsecureEnvironment = "MATTERMOST_INSECURE"

	MattermostPerPage    int = 1000
	MattermostMaxPerPage int = 200

	MattermostSinceChunkThreshold int = 1000
	MattermostSinceChunkLimit     int = 20
	MattermostCollapsedThreads        = false

	MattermostEmptyEntityTag = ""

	MattermostSocketWait = 5 * time.Second

	MattermostKeepAlive  = 25 * time.Second
	MattermostPingWait   = time.Second

	MattermostPostField     = "post"
	MattermostReactionField = "reaction"

	MattermostDoneReaction      = "white_check_mark"
	MattermostProgressReaction  = "construction"
	MattermostWaitingReaction   = "hourglass_flowing_sand"
	MattermostForwardedReaction = "repeat"
	MattermostThreadReaction    = "thread"

	TelegramTokenEnvironment    = "TELEGRAM_TOKEN"
	TelegramChannelEnvironment  = "TELEGRAM_CHANNEL"
	TelegramDatabaseEnvironment = "TELEGRAM_DATABASE"
)

var (
	MattermostFormat = constant.ExtendedColorFormat.Copy()
	TelegramFormat   = constant.ColorFormat.Copy()
)
