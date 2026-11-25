package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"regexp"
)

var Identity = identity.New(
	"gographd",
	"Microsoft 365 bridge through Microsoft Graph",
	"gographd",
).WithInstructions(
	"Microsoft Teams through Microsoft Graph, as the signed-in user. list_teams and list_channels find a channel; get_channel_history and get_thread_replies read it; my_chats and get_chat_history read one-to-one, group and meeting chats; send_chat_message posts to a chat. Teams and channels take a name or an identifier. Sign-in happens once in a terminal with gograph login; until then every tool answers that sign-in is required.",
)

const (
	ListTeams         = "list_teams"
	ListChannels      = "list_channels"
	GetChannelHistory = "get_channel_history"
	GetThreadReplies  = "get_thread_replies"
	MyChats           = "my_chats"
	GetChatHistory    = "get_chat_history"
	SendChatMessage   = "send_chat_message"
)

const GraphPeer = "microsoft graph"
const ParentStep = ".."

var (
	ChatPattern    = regexp.MustCompile(`^[0-9A-Za-z:@._-]+$`)
	MessagePattern = regexp.MustCompile(`^[0-9]+$`)
)
