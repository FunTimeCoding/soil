package response

type RuntimeDependency struct {
	FullName         string `json:"full_name"`
	Version          string `json:"version"`
	Revision         int    `json:"revision,omitzero"`
	BottleRebuild    int    `json:"bottle_rebuild,omitzero"`
	PackageVersion   string `json:"pkg_version,omitempty"`
	DeclaredDirectly bool   `json:"declared_directly"`
}
