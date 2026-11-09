package string_constant

func collectConstants(packageDirectory string) map[string][]KnownConstant {
	result := make(map[string][]KnownConstant)
	collectConstantFile(result, packageDirectory, "")
	collectConstantDirectory(result, packageDirectory, "constant")
	collectParents(result, packageDirectory)

	return result
}
