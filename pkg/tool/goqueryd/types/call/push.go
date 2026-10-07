package call

type Push struct {
	Collection string
	Name       string
	Body       string
	Metadata   map[string][]string
}
