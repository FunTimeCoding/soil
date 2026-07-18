package argument

type RenamePackageClause struct {
	PackagePath string `json:"package_path"`
	NewName     string `json:"new_name"`
}
