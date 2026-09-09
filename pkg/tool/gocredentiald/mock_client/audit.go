package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/audit_report"

func (c *Client) Audit(_ int) *audit_report.Report {
	return audit_report.New()
}
