package rule

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (r *Rule) formatType() string {
	if r.RawAlert != nil {
		return constant.AlertType
	}

	if r.RawRecord != nil {
		return constant.RecordType
	}

	return constant.UnknownType
}
