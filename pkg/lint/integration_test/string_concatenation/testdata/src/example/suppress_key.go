package example

func SuppressKey(a, b string) string {
	return a + b // goanalyze:ignore string_concat
}
