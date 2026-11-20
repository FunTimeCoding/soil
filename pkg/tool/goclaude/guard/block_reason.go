package guard

func BlockReason(
	system string,
	command string,
) string {
	return GrantedBlockReason(system, command, nil)
}
