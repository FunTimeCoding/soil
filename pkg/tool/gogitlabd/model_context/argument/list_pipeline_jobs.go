package argument

type ListPipelineJobs struct {
	Project  string `json:"project"`
	Pipeline int64  `json:"pipeline"`
	Raw      bool   `json:"raw"`
}
