package segment_span

func New(
	start int,
	end int,
	lower string,
) *Span {
	return &Span{Start: start, End: end, Lower: lower}
}
