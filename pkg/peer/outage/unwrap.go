package outage

func (o *OutageError) Unwrap() error {
	return o.cause
}
