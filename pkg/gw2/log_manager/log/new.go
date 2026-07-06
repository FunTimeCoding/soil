package log

import (
	"github.com/funtimecoding/soil/pkg/gw2/log_manager"
	"strings"
)

func New(l *log_manager.Log) *Log {
	var accounts []string

	for _, player := range l.Players {
		accounts = append(
			accounts,
			strings.TrimPrefix(player.AccountName, ":"),
		)
	}

	return &Log{
		Time:     timeFromName(l.FileName),
		Accounts: accounts,
		Raw:      l,
	}
}
