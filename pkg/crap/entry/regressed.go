package entry

func (e *Entry) Regressed(tolerance float64) bool {
	return !e.IsNew() && e.Delta() > tolerance
}
