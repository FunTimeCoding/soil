package fill

func New(
	tab string,
	uid string,
	value string,
) *Fill {
	return &Fill{Tab: tab, UID: uid, Value: value, Snapshot: true}
}
