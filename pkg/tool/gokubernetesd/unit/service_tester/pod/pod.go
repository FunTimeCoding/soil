package pod

type Pod struct {
	Namespace string
	Name      string
	Phase     string
	Restarts  int64
	Labels    map[string]string
}
