package inspection

func New(instance string) *Inspection {
	return &Inspection{Instance: instance}
}
