package log_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
)

func enemies(f fight.Fight) string {
	if !f.Enriched {
		return "-"
	}

	return fmt.Sprintf("%d", f.EnemyCount)
}
