package target

func Used(version string) string {
	return version
}

func Passed(version string) string {
	return version
}

func Forward(version string) string {
	return Passed(version)
}

func Valued(version string) string {
	return ""
}

func Called(version string) string {
	return ""
}

func Local(version string) string {
	return ""
}
