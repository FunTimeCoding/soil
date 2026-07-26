package constant

const (
	SaltHostEnvironment           = "SALT_HOST"
	SaltPortEnvironment           = "SALT_PORT"
	SaltUserEnvironment           = "SALT_USER"
	SaltPasswordEnvironment       = "SALT_PASSWORD"
	SaltAuthenticationEnvironment = "SALT_AUTHENTICATION"

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
)
