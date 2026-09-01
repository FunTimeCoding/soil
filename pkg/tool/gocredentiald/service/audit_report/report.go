package audit_report

import "github.com/funtimecoding/soil/pkg/tool/gocredentiald/service/credential"

type Report struct {
	Stale         []*credential.Credential
	EmptyUser     []*credential.Credential
	EmptyPassword []*credential.Credential
	Duplicates    []*credential.Credential
}
