package budget

type Budget struct {
	limit    int
	spent    int
	opening  int
	empty    bool
	rendered bool
}
