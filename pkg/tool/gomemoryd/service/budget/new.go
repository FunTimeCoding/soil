package budget

func New(limit int) *Budget {
	return &Budget{limit: limit}
}
