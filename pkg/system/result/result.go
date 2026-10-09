package result

type Result struct {
	OutputString string
	ErrorString  string
	Exit         int
	Error        error
}
