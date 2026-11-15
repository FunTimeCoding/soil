package python_replacement

func New(
	end uint,
	target string,
) *Replacement {
	return &Replacement{End: end, Target: target}
}
