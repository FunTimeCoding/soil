package refusal

func New(reason string) *Refusal {
	return &Refusal{reason: reason}
}
