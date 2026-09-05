package example

func Suppressed(a, b string) string {
	return a + b // goanalyze:ignore
}
