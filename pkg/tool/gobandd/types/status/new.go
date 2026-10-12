package status

func New(instance string) *Status {
	return &Status{Instance: instance}
}
