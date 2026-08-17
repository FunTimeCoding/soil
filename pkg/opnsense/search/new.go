package search

func New(phrase string) *Request {
	return &Request{Current: 1, RowCount: -1, SearchPhrase: phrase}
}
