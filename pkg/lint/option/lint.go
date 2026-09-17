package option

type Lint struct {
	Skips         []string
	Scopes        []string
	Registries    []string
	Configuration string
	Verbose       bool
	Census        bool
	Fix           bool
	Summary       bool
}
