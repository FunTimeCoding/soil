package step

func New(
	host string,
	reason string,
) *Step {
	return &Step{Host: host, Reason: reason}
}
