package option

type Lint struct {
	Skips       []string
	ReflowSkips []string
	Scopes      []string
	Registries  []string
	Verbose     bool
	Census      bool
	Fix         bool
	Summary     bool
	Metadata    bool
}
