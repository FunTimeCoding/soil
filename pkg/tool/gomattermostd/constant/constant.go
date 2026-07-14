package constant

import "github.com/funtimecoding/soil/pkg/identity"

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
)
