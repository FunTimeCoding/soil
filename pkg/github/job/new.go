package job

import "github.com/google/go-github/v90/github"

func New(v *github.WorkflowJob) *Job {
	return &Job{
		Identifier: v.GetID(),
		Name:       v.GetName(),
		Status:     v.GetStatus(),
		Conclusion: v.GetConclusion(),
		Hash:       v.GetHeadSHA(),
		CreatedAt:  v.CreatedAt.Time,
		Raw:        v,
	}
}
