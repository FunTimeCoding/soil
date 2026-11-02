package load_failure

func New(
	packagePath string,
	cause error,
) *LoadFailureError {
	return &LoadFailureError{Package: packagePath, Cause: cause}
}
