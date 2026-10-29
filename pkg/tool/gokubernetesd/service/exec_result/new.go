package exec_result

func New(
	stdout string,
	stderr string,
	failed error,
) *Result {
	return &Result{Stdout: stdout, Stderr: stderr, Failed: failed}
}
