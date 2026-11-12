package edit

func New(
	end uint,
	target string,
) *Edit {
	return &Edit{End: end, Target: target}
}
