package string_constant

func ConstantFiles(directory string) []string {
	return ListedConstantFiles(directory, DirectorySources)
}
