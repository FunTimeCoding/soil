package certificate

func Key() string {
	return resolve(KeyFileEnvironment, KeyName)
}
