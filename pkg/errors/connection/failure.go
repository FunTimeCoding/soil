package connection

type Failure struct {
	Kind   string
	Host   string
	Path   string
	Reason string
	class  error
}
