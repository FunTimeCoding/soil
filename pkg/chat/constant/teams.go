package constant

const (
	TeamsPath                = "/v1.0/me/joinedTeams"
	TeamsChannelsPath        = "/v1.0/teams/%s/channels"
	TeamsChannelMessagesPath = "/v1.0/teams/%s/channels/%s/messages"
	TeamsRepliesPath         = "/v1.0/teams/%s/channels/%s/messages/%s/replies"
	TeamsChatsPath           = "/v1.0/me/chats"
	TeamsChatMessagesPath    = "/v1.0/chats/%s/messages"
	TeamsChatMembersPath     = "/v1.0/chats/%s/members"
	TeamsCreateChatPath      = "/v1.0/chats"
	TeamsUsersPath           = "/v1.0/users"
	TeamsUserPath            = "/v1.0/users/%s"
	TeamsUserBinding         = "https://graph.microsoft.com/v1.0/users('%s')"
	TeamsUserMatch           = "startswith(displayName,'%[1]s') or startswith(mail,'%[1]s') or startswith(userPrincipalName,'%[1]s')"
	TeamsUserFields          = "id,displayName,mail,userPrincipalName,jobTitle"
	TeamsMemberType          = "#microsoft.graph.aadUserConversationMember"
	TeamsOwnerRole           = "owner"
	TeamsDirectChat          = "oneOnOne"
	TeamsGroupChat           = "group"
	TeamsRepliesExpansion    = "replies"
	TeamsChatOrder           = "lastMessagePreview/createdDateTime desc"
	TeamsMessageOrder        = "createdDateTime desc"
	TeamsHypertextContent    = "html"
	TeamsUserIdentity        = "aadUser"
	TeamsDefaultLimit        = 20
)
