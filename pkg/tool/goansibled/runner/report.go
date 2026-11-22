package runner

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/provision/types/host_result"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/types/callback_report"
)

func report(output string) []*host_result.Result {
	var r callback_report.Report

	if e := json.Unmarshal([]byte(output), &r); e != nil {
		return nil
	}

	return r.Results()
}
