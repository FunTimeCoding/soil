package host_result

func New(
	host string,
	status string,
	returnCode int,
	stdout string,
	stderr string,
) *Result {
	return &Result{
		Host:       host,
		Status:     status,
		ReturnCode: returnCode,
		Stdout:     stdout,
		Stderr:     stderr,
	}
}
