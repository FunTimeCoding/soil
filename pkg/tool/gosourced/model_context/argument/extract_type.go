package argument

type ExtractType struct {
	PackagePath       string `json:"package_path"`
	Type              string `json:"type"`
	TargetPackagePath string `json:"target_package_path"`
	TargetFile        string `json:"target_file"`
	Create            bool   `json:"create"`
}
