package target

func Mount(
	name string,
	version string,
	port int,
) string {
	return name + string(rune(port))
}
