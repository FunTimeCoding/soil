package constant

const (
	TeamsPath                = "/v1.0/me/joinedTeams"
	TeamsChannelsPath        = "/v1.0/teams/%s/channels"
	TeamsChannelMessagesPath = "/v1.0/teams/%s/channels/%s/messages"
	TeamsRepliesPath         = "/v1.0/teams/%s/channels/%s/messages/%s/replies"
	TeamsChatsPath           = "/v1.0/me/chats"
	TeamsChatMessagesPath    = "/v1.0/chats/%s/messages"
	TeamsChatMembersPath     = "/v1.0/chats/%s/members"
	TeamsRepliesExpansion    = "replies"
	TeamsChatOrder           = "lastMessagePreview/createdDateTime desc"
	TeamsMessageOrder        = "createdDateTime desc"
	TeamsHypertextContent    = "html"
	TeamsUserIdentity        = "aadUser"
	TeamsDefaultLimit        = 20
)
