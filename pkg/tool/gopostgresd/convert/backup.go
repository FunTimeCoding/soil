package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/types/backup"
)

func Backup(b *backup.Backup) server.Backup {
	result := server.Backup{
		Instance: b.Instance,
		Name:     b.Name,
		Time:     b.Time,
		Size:     b.Size,
	}

	if b.Checksum != "" {
		result.Checksum = &b.Checksum
		result.Tables = &b.Tables
	}

	return result
}
