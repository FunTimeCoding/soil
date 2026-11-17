package deletion_result

func New(
	identifier string,
	name string,
) *Result {
	return &Result{Identifier: identifier, Name: name}
}
