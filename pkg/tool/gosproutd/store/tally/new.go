package tally

func New(session string) *Tally {
	return &Tally{Session: session}
}
