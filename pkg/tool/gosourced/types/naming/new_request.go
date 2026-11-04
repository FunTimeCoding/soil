package naming

func NewRequest(
	packagePath string,
	functions []*Function,
	typeSource string,
	name string,
	imports []string,
) *Request {
	return &Request{
		PackagePath: packagePath,
		Functions:   functions,
		Type:        typeSource,
		Name:        name,
		Imports:     imports,
	}
}
