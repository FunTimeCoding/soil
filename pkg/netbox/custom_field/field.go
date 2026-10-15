package custom_field

import "github.com/funtimecoding/soil/pkg/netbox/custom_field_record"

type Field struct {
	Identifier  int32
	Name        string
	Label       string
	Kind        string
	ObjectTypes []string
	ChoiceSet   string
	Raw         *custom_field_record.Record
}
