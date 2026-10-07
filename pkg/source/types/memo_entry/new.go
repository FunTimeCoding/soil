package memo_entry

func New(
	key string,
	value any,
) *Entry {
	return &Entry{Key: key, Value: value}
}
