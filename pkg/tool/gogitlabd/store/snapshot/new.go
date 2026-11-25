package snapshot

import (
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
	"slices"
	"strings"
)

func New(
	k *watched_request.Key,
	d *merge_request_detail.Detail,
	approvers []string,
) *Snapshot {
	v := d.Raw
	result := &Snapshot{
		Instance:    k.Instance,
		Project:     k.Project,
		Request:     k.Request,
		State:       v.State,
		Draft:       v.Draft,
		MergeStatus: v.DetailedMergeStatus,
		Commit:      v.SHA,
		Notes:       v.UserNotesCount,
		Approvers:   strings.Join(approvers, ","),
	}

	if v.UpdatedAt != nil {
		result.Update = *v.UpdatedAt
	}

	if v.HeadPipeline != nil {
		result.Pipeline = v.HeadPipeline.ID
		result.PipelineStatus = v.HeadPipeline.Status
	}

	if slices.Contains(constant.TransientMergeStatus, result.MergeStatus) {
		result.MergeStatus = ""
	}

	return result
}
