package option

type Fix struct {
	Root         string
	Patterns     []string
	Diff         bool
	Full         bool
	Ours         []string
	Index        string
	Replacing    []string
	BeforeCommit func()
}
