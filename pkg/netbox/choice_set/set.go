package choice_set

import "github.com/funtimecoding/soil/pkg/netbox/choice_set_record"

type Set struct {
	Identifier int32
	Name       string
	Choices    []string
	Raw        *choice_set_record.Record
}
