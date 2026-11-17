package byte_limit

type Limit struct {
	limit    int
	used     int
	opening  int
	empty    bool
	rendered bool
}
