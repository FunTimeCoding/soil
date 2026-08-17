package string_constant

func collectConstants(packageDirectory string) map[string][]knownConstant {
	result := make(map[string][]knownConstant)
	collectFromConstantFile(result, packageDirectory, "")
	collectFromConstantDirectory(result, packageDirectory, "constant")
	collectFromParents(result, packageDirectory)

	return result
}
