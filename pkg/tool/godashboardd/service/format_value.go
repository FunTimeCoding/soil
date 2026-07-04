package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/constant"
	"github.com/prometheus/common/model"
	"strconv"
)

func formatValue(v model.Value) string {
	vector, okay := v.(model.Vector)

	if !okay || len(vector) == 0 {
		return constant.PendingValue
	}

	return strconv.FormatInt(int64(vector[0].Value), 10)
}
