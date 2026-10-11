package packager

func NewCustom(options ...Option) *Packager {
	result := &Packager{}

	for _, o := range options {
		o(result)
	}

	return result
}
