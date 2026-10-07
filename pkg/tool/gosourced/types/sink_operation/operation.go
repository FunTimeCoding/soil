package sink_operation

type Operation struct {
	Kind    string
	Path    string
	Target  string
	Content []byte
}
