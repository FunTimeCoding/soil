package policy

type Policy struct {
	RequireCountry      bool
	RequireProvince     bool
	RequireOrganization bool
	InheritIssuer       bool
}
