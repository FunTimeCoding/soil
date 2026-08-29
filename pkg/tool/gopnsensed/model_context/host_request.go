package model_context

import (
	"github.com/funtimecoding/soil/pkg/opnsense/request"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func hostRequest(r mcp.CallToolRequest) *request.Host {
	result := request.New()
	result.Host = r.GetString(constant.ParameterHost, "")
	result.Domain = r.GetString(constant.ParameterDomain, "")
	result.Address = r.GetString(constant.ParameterAddress, "")
	result.HardwareAddress = r.GetString(constant.ParameterHardwareAddress, "")
	result.ClientIdentifier = r.GetString(
		constant.ParameterClientIdentifier,
		"",
	)
	result.Description = r.GetString(constant.ParameterDescription, "")

	return result
}
