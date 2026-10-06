package stalled

func (*Stalled) Error() string {
	return "i/o timeout"
}
