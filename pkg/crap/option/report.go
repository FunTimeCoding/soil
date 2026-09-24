package option

type Report struct {
	Root           string
	Patterns       []string
	Profile        string
	Mutation       string
	Baseline       string
	Missing        string
	Threshold      float64
	Minimum        float64
	Tolerance      float64
	Top            int
	FailAbove      bool
	FailRegression bool
	IgnoreCovered  bool
	All            bool
	Notation       bool
	Verbose        bool
}
