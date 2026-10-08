package string_constant

func collectConstants(packageDirectory string) map[string][]KnownConstant {
	result := make(map[string][]KnownConstant)
	collectFromConstantFile(result, packageDirectory, "")
	collectFromConstantDirectory(result, packageDirectory, "constant")
	collectFromParents(result, packageDirectory)

	return result
}
