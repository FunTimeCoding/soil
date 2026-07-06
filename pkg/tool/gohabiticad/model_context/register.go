package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetTasks,
			mcp.WithDescription(
				"List Habitica tasks. Optionally filter by type: habits, dailys, todos, rewards.",
			),
			mcp.WithString(
				constant.TaskType,
				mcp.Description("Task type filter: habits, dailys, todos, rewards"),
			),
		),
		s.getTasks,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTask,
			mcp.WithDescription("Create a new Habitica task"),
			mcp.WithString(
				constant.TaskType,
				mcp.Required(),
				mcp.Description("Task type: habit, daily, todo, reward"),
			),
			mcp.WithString(
				constant.Text,
				mcp.Required(),
				mcp.Description("Task title"),
			),
			mcp.WithString(
				constant.Notes,
				mcp.Description("Task notes"),
			),
		),
		s.createTask,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ScoreTask,
			mcp.WithDescription(
				"Score a Habitica task. For habits use direction up or down. For dailies and todos use up to complete.",
			),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Task ID"),
			),
			mcp.WithString(
				constant.Direction,
				mcp.Description("Score direction: up (default) or down"),
			),
		),
		s.scoreTask,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetTags,
			mcp.WithDescription("List all Habitica tags"),
		),
		s.getTags,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetStats,
			mcp.WithDescription(
				"Get Habitica user stats: HP, MP, XP, gold, level, class, unspent points, and allocated stats.",
			),
		),
		s.getStats,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AllocateStat,
			mcp.WithDescription(
				"Allocate one unspent stat point. Returns updated stats.",
			),
			mcp.WithString(
				constant.Stat,
				mcp.Required(),
				mcp.Description("Stat to allocate to: str, con, int, or per"),
			),
		),
		s.allocateStat,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetGear,
			mcp.WithDescription(
				"List equipped gear slots and all owned gear keys.",
			),
		),
		s.getGear,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.EquipGear,
			mcp.WithDescription(
				"Equip or unequip a gear item (toggles). Returns updated gear state.",
			),
			mcp.WithString(
				constant.Key,
				mcp.Required(),
				mcp.Description("Gear key to equip or unequip"),
			),
		),
		s.equipGear,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Cron,
			mcp.WithDescription(
				"Check and run the daily rollover if needed. Applies damage for incomplete dailies and resets daily tasks. Call this before any other Habitica tool to ensure the day state is current. Returns stat changes if rollover ran.",
			),
		),
		s.cron,
	)
}
