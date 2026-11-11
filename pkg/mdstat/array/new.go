package array

func New(
	name string,
	state string,
	level string,
) *Array {
	return &Array{Name: name, State: state, Level: level}
}
