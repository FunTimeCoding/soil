package face

import (
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/audit_report"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/entry_detail"
)

type CredentialSource interface {
	List() []*credential.Credential
	Search(query string) []*credential.Credential
	Get(identifier string) *entry_detail.Detail
	Reveal(identifier string) (string, bool)
	Audit(staleYears int) *audit_report.Report
	LoadGroup(name string) map[string]string
	Create(
		groupPath string,
		title string,
		fields map[string]string,
	) (string, error)
	Update(
		identifier string,
		fields map[string]string,
	) error
	Move(
		identifier string,
		groupPath string,
	) error
	Delete(identifier string) error
}
