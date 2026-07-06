package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/raid"
	"strings"
)

func fightMap(f raid.Fight) string {
	if strings.Contains(f.MapName, "Green Alpine") {
		return "GBL"
	}

	if strings.Contains(f.MapName, "Blue Alpine") {
		return "BBL"
	}

	if strings.Contains(f.MapName, "Eternal Battlegrounds") {
		return "EBG"
	}

	if strings.Contains(f.MapName, "Red Desert") {
		return "RBL"
	}

	if f.MapName != "" {
		return f.MapName
	}

	return fmt.Sprintf("%d", f.MapID)
}
