package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

type Client interface {
	Close()
	CreateBackup(
		x context.Context,
		body client.CreateBackupJSONRequestBody,
	) (*client.Backup, error)
	DeleteBackup(
		x context.Context,
		name string,
		params *client.DeleteBackupParams,
	) error
	DescribeTable(
		x context.Context,
		table string,
		params *client.DescribeTableParams,
	) (*client.QueryResult, error)
	Explain(
		x context.Context,
		body client.ExplainJSONRequestBody,
	) (*client.QueryResult, error)
	ListIndexes(
		x context.Context,
		params *client.ListIndexesParams,
	) (*client.QueryResult, error)
	ListBackups(
		x context.Context,
		params *client.ListBackupsParams,
	) (*client.BackupList, error)
	ListInstances(x context.Context) ([]client.Instance, error)
	ListSchemas(
		x context.Context,
		params *client.ListSchemasParams,
	) (*client.QueryResult, error)
	ListTables(
		x context.Context,
		params *client.ListTablesParams,
	) (*client.QueryResult, error)
	Query(
		x context.Context,
		body client.QueryJSONRequestBody,
	) (*client.QueryResult, error)
	Reachable(x context.Context) error
	RestoreBackup(
		x context.Context,
		name string,
		body client.RestoreBackupJSONRequestBody,
	) (*client.RestoreResult, error)
	TableSizes(
		x context.Context,
		params *client.TableSizesParams,
	) (*client.QueryResult, error)
}
