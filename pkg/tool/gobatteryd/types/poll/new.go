package poll

func New() *Poll {
	return &Poll{Reachable: map[string]bool{}}
}
