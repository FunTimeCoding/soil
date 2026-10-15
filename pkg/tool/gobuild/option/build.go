package option

type Build struct {
	Names      []string
	Module     string
	Version    string
	Native     bool
	Tags       string
	Main       string
	Output     string
	Targets    []string
	Install    bool
}
