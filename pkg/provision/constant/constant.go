package constant

import "time"

const (
	AnsibleInventoryEnvironment = "ANSIBLE_INVENTORY"

	DownstreamClient  = "downstream"
	DownstreamTool    = "trigger"
	DownstreamUpdate  = "update"
	DownstreamChanges = "changes"

	PackerDirectory       = "packer"
	PackerWebDirectory    = "web"
	PackerOutputDirectory = "output"

	RunnerSyncInterval  = 5 * time.Minute
	RunnerApplyInterval = 30 * time.Minute
	RunnerBranch        = "main"
	RunnerRemoteBranch  = "origin/main"

	RunnerTriggerTimer  = "timer"
	RunnerTriggerManual = "manual"

	RunnerStatus = "status"
	RunnerLocal  = "local"
	RunnerRemote = "remote"
	RunnerPath   = "path"

	RunnerIndexLock = ".git/index.lock"

	RunnerQuarantineSuffix = ".quarantine."
	RunnerQuarantineFormat = "20060102T150405Z"

	RunnerHealThreshold = 3
	RunnerStagePattern  = "provision-clone-"
	RunnerStageSuffix   = ".stage"

	RunnerConsecutive = "consecutive"
	RunnerError       = "error"

	SaltHostEnvironment           = "SALT_HOST"
	SaltPortEnvironment           = "SALT_PORT"
	SaltUserEnvironment           = "SALT_USER"
	SaltPasswordEnvironment       = "SALT_PASSWORD"
	SaltAuthenticationEnvironment = "SALT_AUTHENTICATION"
	SaltInsecureEnvironment       = "SALT_INSECURE"

	SaltRun       = "cmd.run"
	SaltHighstate = "state.highstate"

	SaltTokenHeader = "X-Auth-Token"

	SaltLoginPath   = "login"
	SaltKeysPath    = "keys"
	SaltMinionsPath = "minions"
	SaltJobsPath    = "jobs"

	SaltLocalClient      = "local"
	SaltLocalAsyncClient = "local_async"
	SaltWheelClient      = "wheel"

	SaltGlobTarget = "glob"

	SaltKeyAccept = "key.accept"
	SaltKeyDelete = "key.delete"

	StoreStatusRunning = "running"
	StoreStatusSuccess = "success"
	StoreStatusError   = "error"
	StoreRetentionDays = 14
	StoreRetentionAge  = StoreRetentionDays * 24 * time.Hour
)
