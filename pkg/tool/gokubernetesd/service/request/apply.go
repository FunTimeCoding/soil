package request

type Apply struct {
	Manifest  string
	Namespace string
	Override  bool
	DryRun    bool
}
