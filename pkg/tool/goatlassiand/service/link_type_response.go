package service

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/link_type"

type linkTypeResponse struct {
	IssueLinkTypes []link_type.Type `json:"issueLinkTypes"`
}
