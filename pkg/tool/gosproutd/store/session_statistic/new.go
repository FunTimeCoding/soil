package session_statistic

func New(session string) *Statistic {
	return &Statistic{Session: session}
}
