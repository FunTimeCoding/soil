package argument

type MovePackage struct {
	PackagePath       string `json:"package_path"`
	TargetPackagePath string `json:"target_package_path"`
}
