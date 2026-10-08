package guard

func newEdit(
	end uint,
	target string,
) *Edit {
	return &Edit{end: end, target: target}
}
