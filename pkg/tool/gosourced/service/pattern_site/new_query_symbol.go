package pattern_site

func NewQuerySymbol(
	name string,
	receiver string,
) *QuerySymbol {
	return &QuerySymbol{Name: name, Receiver: receiver}
}
