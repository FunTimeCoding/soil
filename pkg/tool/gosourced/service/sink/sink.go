package sink

type Sink struct {
	root       string
	operations []*operation
	dropped    []string
}
