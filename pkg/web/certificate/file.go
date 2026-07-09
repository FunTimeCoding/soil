package certificate

func File() string {
	return resolve(FileEnvironment, FileName)
}
