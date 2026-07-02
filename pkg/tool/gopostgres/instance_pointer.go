package gopostgres

func instancePointer(instance string) *string {
	if instance == "" {
		return nil
	}

	return &instance
}
