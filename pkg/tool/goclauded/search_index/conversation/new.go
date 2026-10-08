package conversation

func New(
	session string,
	latest string,
	count int,
) *Conversation {
	return &Conversation{Session: session, Latest: latest, Count: count}
}
