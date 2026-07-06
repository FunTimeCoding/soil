package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid"
)

func fightEnemies(f raid.Fight) string {
	if !f.Enriched {
		return "-"
	}

	return fmt.Sprintf("%d", f.EnemyCount)
}
