package exec_result

type Result struct {
	Stdout string
	Stderr string
	Failed error
}
