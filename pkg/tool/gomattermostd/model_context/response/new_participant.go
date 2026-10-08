package response

func NewParticipant(username string) *Participant {
	return &Participant{Username: username}
}
