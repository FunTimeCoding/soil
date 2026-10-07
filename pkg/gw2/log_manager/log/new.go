package log

import (
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/response"
	"strings"
)

func New(l *response.Log) *Log {
	var accounts []string

	for _, p := range l.Players {
		accounts = append(accounts, strings.TrimPrefix(p.AccountName, ":"))
	}

	return &Log{
		Time:     timeFromName(l.FileName),
		Accounts: accounts,
		Raw:      l,
	}
}
