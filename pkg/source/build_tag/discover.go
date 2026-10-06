package build_tag

func Discover(directory string) []string {
	return Union(Files(directory))
}
