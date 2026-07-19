package argument

type MoveSymbols struct {
	PackagePath           string   `json:"package_path"`
	Symbols               []string `json:"symbols"`
	File                  string   `json:"file"`
	TargetPackagePath     string   `json:"target_package_path"`
	TargetFile            string   `json:"target_file"`
	Create                bool     `json:"create"`
	QualifyBackReferences bool     `json:"qualify_back_references"`
}
