package destination

func Chat(identifier string) *Destination {
	return &Destination{Chat: identifier}
}
