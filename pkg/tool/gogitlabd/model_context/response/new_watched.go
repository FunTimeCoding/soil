package response

import (
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
)

func NewWatched(
	v *subscription.Subscription,
	d *merge_request_detail.Detail,
) *Watched {
	result := &Watched{
		Reference:   v.Reference,
		Alias:       v.Alias,
		Title:       d.Title,
		State:       d.State,
		Draft:       d.Raw.Draft,
		MergeStatus: d.Raw.DetailedMergeStatus,
		Delivery:    string(v.DeliveryMode()),
	}

	if p := d.Raw.HeadPipeline; p != nil {
		result.PipelineStatus = p.Status
	}

	return result
}
