package response

func NewThread(
	postIdentifier string,
	replyCount int64,
	unreadReplies int64,
	unreadMentions int64,
	lastReplyAt string,
) *Thread {
	return &Thread{
		PostIdentifier: postIdentifier,
		ReplyCount:     replyCount,
		UnreadReplies:  unreadReplies,
		UnreadMentions: unreadMentions,
		LastReplyAt:    lastReplyAt,
	}
}
