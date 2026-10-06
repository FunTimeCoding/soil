package argument

type IntroduceConstructor struct {
	PackagePath         string   `json:"package_path"`
	Type                string   `json:"type"`
	Parameters          []string `json:"parameters"`
	ParametersFromSites bool     `json:"parameters_from_sites"`
	AssignRest          bool     `json:"assign_rest"`
	DryRun              bool     `json:"dry_run"`
}
