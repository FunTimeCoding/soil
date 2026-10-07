package link_type_response

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/link_type"

type Response struct {
	IssueLinkTypes []link_type.Type `json:"issueLinkTypes"`
}
