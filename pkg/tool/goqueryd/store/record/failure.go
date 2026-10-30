package record

type Failure struct {
	Collection string
	Path       string
	Body       string
	Sequence   int
	Reason     string
}
