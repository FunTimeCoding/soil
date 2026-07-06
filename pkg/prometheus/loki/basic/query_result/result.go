package query_result

import "github.com/funtimecoding/soil/pkg/prometheus/loki/basic/response"

type Result struct {
	ResultType string             `json:"resultType"`
	Result     []response.Result  `json:"result"`
	Stats      response.Statistic `json:"stats"`
}
