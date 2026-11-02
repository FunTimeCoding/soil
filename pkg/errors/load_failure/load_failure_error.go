package load_failure

type LoadFailureError struct {
	Package string
	Cause   error
	Settled bool
}
