package field_changer

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/field_changer/severity"

type Changer struct {
	name     map[string]string
	severity []*severity.Severity
}
