package option

type Fix struct {
	Root         string
	Patterns     []string
	Diff         bool
	Full         bool
	Realias      bool
	Index        string
	Replacing    []string
	BeforeCommit func()
}
