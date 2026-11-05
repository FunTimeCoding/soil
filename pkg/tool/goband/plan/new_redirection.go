package plan

func NewRedirection(
	instance string,
	pending []string,
) *Redirection {
	return &Redirection{Instance: instance, Pending: pending}
}
