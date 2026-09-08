package utilization

func ReadCredential() *Credential {
	if !Supported() {
		return nil
	}

	return ParseCredential(rawCredential())
}
