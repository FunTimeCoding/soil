package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getGroup(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier := int64(q.GetFloat(constant.MemoryIdentifier, 0))

	if identifier == 0 {
		return response.Fail("memory_id is required")
	}

	parent, children, e := s.service.GetMemoryGroup(identifier)

	if e != nil {
		return s.captureFail(e, "load memory group")
	}

	members := map[int64]bool{parent.Identifier: true}
	identifiers := []int64{parent.Identifier}

	for i := range children {
		members[children[i].Identifier] = true
		identifiers = append(identifiers, children[i].Identifier)
	}

	edges, f := s.service.ListRelationsAmong(identifiers)

	if f != nil {
		return s.captureFail(f, "load group relations")
	}

	relations := groupRelations(edges, members)
	payload := map[string]any{}

	if len(relations) > 0 {
		payload["relations"] = relations
	}

	if q.GetBool(constant.Detail, false) {
		payload["parent"] = parent
		payload["children"] = children

		return response.Success(notation.MarshalIndent(payload))
	}

	payload["parent"] = convert.Memory(parent)
	payload["children"] = convert.Memories(children)

	return response.Success(notation.MarshalIndent(payload))
}
