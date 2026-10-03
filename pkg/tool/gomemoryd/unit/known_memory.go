package unit

func knownMemory(
	scope string,
	name string,
) bool {
	return scope == "default" && name == "terse communication"
}
