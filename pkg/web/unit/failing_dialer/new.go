package failing_dialer

func New(failure error) *Dialer {
	return &Dialer{failure: failure}
}
