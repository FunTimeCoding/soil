package argument

type RenamePackage struct {
	PackagePath string `json:"package_path"`
	NewName     string `json:"new_name"`
}
