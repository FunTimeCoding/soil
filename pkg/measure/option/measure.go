package option

type Measure struct {
	Paths     []string
	Skips     []string
	Languages []string
	ByFile    bool
	Notation  bool
	Verbose   bool
}
