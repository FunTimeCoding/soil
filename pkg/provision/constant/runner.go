package constant

import "time"

const (
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

	RunnerConsecutive = "consecutive"
	RunnerError       = "error"
)
