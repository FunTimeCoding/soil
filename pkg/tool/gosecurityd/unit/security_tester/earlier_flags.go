package security_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/migrate"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store"
	"gorm.io/gorm"
	"path/filepath"
)

func (o *Tester) EarlierFlags(subject string) (*store.Store, *gorm.DB) {
	o.t.Helper()
	m := lite.New(
		logger.New(o.t.Context()),
		filepath.Join(o.t.TempDir(), "earlier.sqlite"),
	)
	o.t.Cleanup(func() { relational.CloseMapper(m) })

	for _, statement := range []string{
		"CREATE TABLE `security_flags` (`id` integer PRIMARY KEY AUTOINCREMENT,`vulnerability` text NOT NULL,`note` text,`severity` text NOT NULL DEFAULT \"\",`score` real NOT NULL DEFAULT 0,`mitigated` numeric NOT NULL DEFAULT false,`mitigation` text,`created_at` datetime NOT NULL,`updated_at` datetime NOT NULL)",
		"CREATE INDEX `idx_security_flags_mitigated` ON `security_flags`(`mitigated`)",
		"CREATE UNIQUE INDEX `idx_security_flags_vulnerability` ON `security_flags`(`vulnerability`)",
		"CREATE UNIQUE INDEX `idx_security_flags_subject` ON `security_flags`(`vulnerability`)",
	} {
		assert.FatalOnError(o.t, m.Exec(statement).Error)
	}

	assert.FatalOnError(
		o.t,
		m.Exec(
			"INSERT INTO `security_flags` (`vulnerability`, `note`, `created_at`, `updated_at`) VALUES (?, 'kept', datetime('now'), datetime('now'))",
			subject,
		).Error,
	)
	migrate.AutoMigrate(m)

	return store.New(m), m
}
