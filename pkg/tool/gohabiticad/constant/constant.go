package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gohabiticad",
	"Habitica task gamification bridge",
	"gohabiticad",
).WithInstructions(
	"Habitica task gamification - score tasks, check stats, manage gear. Habits use up/down direction, dailies and todos use up to complete.",
)

const (
	GetTasks     = "get_tasks"
	CreateTask   = "create_task"
	ScoreTask    = "score_task"
	GetTags      = "get_tags"
	GetStats     = "get_stats"
	AllocateStat = "allocate_stat"
	GetGear      = "get_gear"
	EquipGear    = "equip_gear"
	Cron         = "cron"
	Direction    = "direction"
	Key          = "key"
	Notes        = "notes"
	Stat         = "stat"
	TaskType     = "task_type"
	Text         = "text"
)
