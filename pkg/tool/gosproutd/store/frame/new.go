package frame

func New(session string, name string) *Frame {
	return &Frame{Session: session, Name: name}
}
