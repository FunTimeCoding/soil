package job

import (
	"github.com/funtimecoding/soil/pkg/monitor/item/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func New(v *gitlab.Job) *Job {
	return &Job{
		MonitorIdentifier: constant.GoGitLab.Integer64Identifier(v.ID),
		Identifier:        v.ID,
		Name:              v.Name,
		Status:            v.Status,
		Stage:             v.Stage,
		Create:            v.CreatedAt,
		Link:              v.WebURL,
		Raw:               v,
	}
}
