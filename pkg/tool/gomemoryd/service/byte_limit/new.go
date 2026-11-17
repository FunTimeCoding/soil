package byte_limit

func New(limit int) *Limit {
	return &Limit{limit: limit}
}
