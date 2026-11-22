package host_result

type Result struct {
	Host       string
	Status     string
	ReturnCode int
	Stdout     string
	Stderr     string
}
