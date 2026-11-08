package target

func (t *Target) Holds(
	key string,
	value string,
) bool {
	return t.Held(key, value) != ""
}
