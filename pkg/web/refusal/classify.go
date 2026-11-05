package refusal

func (t *Table) Classify(
	status int,
	reason string,
) error {
	class := t.Class(status)

	if class == "" {
		return nil
	}

	return Classed(class, reason)
}
