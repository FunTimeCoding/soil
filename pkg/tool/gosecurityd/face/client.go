package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

type Client interface {
	GetVulnerability(
		x context.Context,
		identifier string,
	) (*client.Vulnerability, error)
	GetDebianStatus(
		x context.Context,
		identifier string,
		p *client.GetDebianStatusParams,
	) (*client.DebianStatus, error)
	ListFlags(
		x context.Context,
		p *client.ListFlagsParams,
	) ([]*client.Flag, error)
	FlagVulnerability(
		x context.Context,
		b client.FlagRequest,
	) (*client.Flag, error)
	MitigateFlag(
		x context.Context,
		identifier string,
		b client.MitigationRequest,
	) (*client.Flag, error)
	PutSystem(
		x context.Context,
		name string,
		b client.SystemSnapshot,
	) error
	ListSystems(
		x context.Context,
		p *client.ListSystemsParams,
	) ([]*client.SystemSummary, error)
	GetSystemFindings(
		x context.Context,
		name string,
	) (*client.SystemDetail, error)
	DeleteSystem(
		x context.Context,
		name string,
	) error
	ListFindings(
		x context.Context,
		p *client.ListFindingsParams,
	) ([]*client.FindingGroup, error)
	GetScoreStatus(x context.Context) (*client.ScoreStatus, error)
	GetAffectedSystems(
		x context.Context,
		identifier string,
	) (*client.AffectedSystems, error)
	ListUpgrades(
		x context.Context,
		p *client.ListUpgradesParams,
	) (*client.UpgradeList, error)
	GetUpgrade(
		x context.Context,
		release string,
		source string,
	) (*client.UpgradeGroup, error)
	ListRemediation(
		x context.Context,
		p *client.ListRemediationParams,
	) (*client.RemediationList, error)
	Reachable(x context.Context) error
	Address() string
	Close()
}
