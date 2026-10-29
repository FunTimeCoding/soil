package failure

func New(message string) *Failure {
	return &Failure{Error: message}
}
