package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/types/backup"
)

func Backups(v []*backup.Backup) server.BackupList {
	result := server.BackupList{Backups: []server.Backup{}}

	for _, b := range v {
		result.Backups = append(result.Backups, Backup(b))
	}

	return result
}
