package refusal

func Status(
	status int,
	reason string,
) error {
	return NewTable().Classify(status, reason)
}
