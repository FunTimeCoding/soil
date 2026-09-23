package session

func New(name string) *Session {
	return &Session{Name: name}
}
