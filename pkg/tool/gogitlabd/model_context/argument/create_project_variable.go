package argument

type CreateProjectVariable struct {
	Project   string `json:"project"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	Protected bool   `json:"protected"`
	Masked    bool   `json:"masked"`
	Expand    bool   `json:"expand"`
}
