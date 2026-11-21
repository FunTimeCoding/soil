package package_group

type Group struct {
	Source          string
	Kind            string
	Installed       string
	Fixed           string
	Candidate       string
	TopScore        float64
	Vulnerabilities []string
}
