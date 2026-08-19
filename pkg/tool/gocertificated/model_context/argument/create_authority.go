package argument

type CreateAuthority struct {
	Name             string   `json:"name"`
	Kind             string   `json:"kind"`
	CommonName       string   `json:"common_name"`
	Country          string   `json:"country"`
	Province         string   `json:"province"`
	Organization     string   `json:"organization"`
	PermittedDomain  []string `json:"permitted_domain"`
	PermittedAddress []string `json:"permitted_address"`
	ValidYear        float64  `json:"valid_year"`
}
