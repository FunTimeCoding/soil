package operation

func New(
	name string,
	force bool,
	state string,
	refusal string,
) *Operation {
	return &Operation{Name: name, Force: force, State: state, Refusal: refusal}
}
