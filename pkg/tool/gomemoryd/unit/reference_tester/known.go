package reference_tester

func Known(
	scope string,
	name string,
) bool {
	return scope == "default" && name == "terse communication"
}
