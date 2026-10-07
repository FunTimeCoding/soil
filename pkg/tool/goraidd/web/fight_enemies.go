package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
)

func fightEnemies(f fight.Fight) string {
	if !f.Enriched {
		return "-"
	}

	return fmt.Sprintf("%d", f.EnemyCount)
}
