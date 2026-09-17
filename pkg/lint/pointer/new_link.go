package pointer

func NewLink(target string) *Candidate {
	return &Candidate{Span: target, Link: true}
}
