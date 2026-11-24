package remediable

type RemediableError struct {
	Name    string
	Message string
	cause   error
}
