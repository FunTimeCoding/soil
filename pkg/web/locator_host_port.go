package web

func LocatorHostPort(s string) string {
	return ParseLocator(s).Host
}
