package lib

func Log(prefix string) string {
	return prefix
}

func Other() string {
	return Log("other")
}
