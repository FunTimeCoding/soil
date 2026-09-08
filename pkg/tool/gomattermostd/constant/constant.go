package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"time"
)

var Identity = identity.New(
	"gomattermostd",
	"Mattermost messaging bridge",
	"gomattermostd",
).WithInstructions(
	"Mattermost messaging - channels, threads, search, direct messages, reactions, file sharing.",
)

const (
	MyChannels        = "my_channels"
	MyThreads         = "my_threads"
	SearchChannels    = "search_channels"
	SearchMessages    = "search_messages"
	SearchUsers       = "search_users"
	DownloadFile      = "download_file"
	UploadFile        = "upload_file"
	ListChannels      = "list_channels"
	GetChannel        = "get_channel"
	GetChannelHistory = "get_channel_history"
	SendDirectMessage = "send_dm"
	PostMessage       = "post_message"
	ReplyToThread     = "reply_to_thread"
	AddReaction       = "add_reaction"
	GetThreadReplies  = "get_thread_replies"
	GetUsers          = "get_users"
	GetUserProfile    = "get_user_profile"
	RunMonitoring     = "run_monitoring"
	SubscribeThread   = "subscribe_thread"
	UnsubscribeThread = "unsubscribe_thread"
	ListSubscriptions = "list_subscriptions"
)

const (
	ParameterRoot     = "root"
	ParameterCallsign = "callsign"
	ParameterAlias    = "alias"
)

const (
	ExcerptLength    = 200
	DigestBudget     = 700
	EnumerateLimit   = 4
	TruncationMarker = "…+more"
	DigestIndent     = "    "
	LabelPrefix      = 8
)

const (
	ReconnectDelay = 2 * time.Second
	DebounceWindow = 30 * time.Second
	PurgeWindow    = 7 * 24 * time.Hour
	PurgeInterval  = time.Hour
)
const (
	HostEnvironment     = "GOMATTERMOST_HOST"
	PortEnvironment     = "GOMATTERMOST_PORT"
	InsecureEnvironment = "GOMATTERMOST_INSECURE"
	TokenEnvironment    = "GOMATTERMOST_TOKEN" // #nosec G101 not a hardcoded secret
)
