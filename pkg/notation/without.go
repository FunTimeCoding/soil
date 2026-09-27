package notation

func Without(
	a any,
	key string,
) any {
	var v any
	MustDecodeBytes(MarshalIndentBytes(a), &v, false)

	return withoutValue(v, key)
}
