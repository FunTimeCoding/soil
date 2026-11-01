package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerJournal() {
	s.server.AddTool(
		mcp.NewTool(
			constant.AddDeviceJournalEntry,
			mcp.WithDescription(
				"Add a journal entry to a device - an append-only timestamped note, visible on the device page in NetBox",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Comments,
				mcp.Required(),
				mcp.Description("Entry text, markdown"),
			),
			mcp.WithString(
				constant.Kind,
				mcp.Description("Entry kind: info, success, warning, or danger"),
			),
		),
		s.addDeviceJournalEntry,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDeviceJournalEntries,
			mcp.WithDescription(
				"List journal entries for a device, newest first",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithNumber(
				generativeConstant.ParameterLimit,
				mcp.Description("Maximum entries to return"),
			),
			mcp.WithNumber(
				generativeConstant.ParameterOffset,
				mcp.Description("Entries to skip"),
			),
		),
		s.listDeviceJournalEntries,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddVirtualJournalEntry,
			mcp.WithDescription(
				"Add a journal entry to a virtual machine - an append-only timestamped note, visible on the VM page in NetBox",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Comments,
				mcp.Required(),
				mcp.Description("Entry text, markdown"),
			),
			mcp.WithString(
				constant.Kind,
				mcp.Description("Entry kind: info, success, warning, or danger"),
			),
		),
		s.addVirtualJournalEntry,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListVirtualJournalEntries,
			mcp.WithDescription(
				"List journal entries for a virtual machine, newest first",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithNumber(
				generativeConstant.ParameterLimit,
				mcp.Description("Maximum entries to return"),
			),
			mcp.WithNumber(
				generativeConstant.ParameterOffset,
				mcp.Description("Entries to skip"),
			),
		),
		s.listVirtualJournalEntries,
	)
}
