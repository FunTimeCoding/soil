package constant

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/identity/paragraph"
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"time"
)

var Identity = identity.New(
	"gopostgresd",
	"Multi-instance PostgreSQL access bridge",
	"gopostgresd",
).WithInstructions(
	"PostgreSQL access.",
	paragraph.New(
		inventoryConstant.MultiInstance,
		inventoryConstant.MultiInstanceInstruction,
	),
)

const (
	Query         = "query"
	Explain       = "explain"
	ListSchemas   = "list_schemas"
	ListTables    = "list_tables"
	DescribeTable = "describe_table"
	ListIndexes   = "list_indexes"
	TableSizes    = "table_sizes"
)

const (
	DefaultSchema = "public"
	InstancePeer  = "postgres"
)

var BackupOperations = []string{
	"createBackup",
	"listBackups",
	"deleteBackup",
	"restoreBackup",
}

const (
	CreateBackup  = "create_backup"
	ListBackups   = "list_backups"
	RestoreBackup = "restore_backup"
	DeleteBackup  = "delete_backup"

	BackupKind              = "backup"
	BackupParameter         = "backup"
	BackupDirectoryArgument = "backup-directory"
	BackupNameLayout        = "2006-01-02T15-04-05"
	BackupExtension         = ".dump"
	ChecksumExtension       = ".sha256"
	ChecksumLine            = "%s  %s\n"
	RestoreSuffix           = "_restore_"
	TableContentMarker      = " TABLE DATA "

	DumpProgram    = "pg_dump"
	RestoreProgram = "pg_restore"

	PasswordEnvironment        = "PGPASSWORD"
	ApplicationNameEnvironment = "PGAPPNAME"
	ApplicationNameParameter   = "application_name"
	RestoreApplicationName     = "gopostgresd restore"

	DumpTimeout         = 30 * time.Minute
	RestoreTimeout      = 60 * time.Minute
	BackupWriteDeadline = RestoreTimeout + time.Minute

	BackupDirectoryMode = 0o750
	BackupFileMode      = 0o640

	BackupNameRefused    = "backup name %q is not a %s timestamp"
	DatabaseTaken        = "database %q exists and is not the instance's own"
	SessionsConnected    = "restore in place refused, other sessions are connected: %s"
	SessionDescription   = "%s (%s@%s since %s)"
	UnnamedApplication   = "(no application name)"
	BackupDirectoryUnset = "backup directory is not configured"
	BackupDeleted        = "backup %s of %s deleted"
)
