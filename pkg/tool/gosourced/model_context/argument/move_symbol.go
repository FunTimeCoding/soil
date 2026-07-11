package argument

type MoveSymbol struct {
	PackagePath       string `json:"package_path"`
	Symbol            string `json:"symbol"`
	TargetPackagePath string `json:"target_package_path"`
	TargetFile        string `json:"target_file"`
	Create            bool   `json:"create"`
}
