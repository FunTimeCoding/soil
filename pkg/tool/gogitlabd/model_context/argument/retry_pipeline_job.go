package argument

type RetryPipelineJob struct {
	Project string `json:"project"`
	Job     int64  `json:"job"`
	Raw     bool   `json:"raw"`
}
