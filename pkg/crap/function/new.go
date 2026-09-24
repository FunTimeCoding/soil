package function

func New(
	packagePath string,
	file string,
	line int,
	name string,
	complexity int,
) *Function {
	return &Function{
		Package:    packagePath,
		File:       file,
		Line:       line,
		Name:       name,
		Complexity: complexity,
	}
}
