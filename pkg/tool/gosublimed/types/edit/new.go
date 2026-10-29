package edit

func New(
	identifier int,
	old string,
	new string,
	all bool,
) *Edit {
	return &Edit{Identifier: identifier, Old: old, New: new, All: all}
}
