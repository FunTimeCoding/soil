package above_below

func Integer(
	i int,
	magnitude int,
	above func(),
	below func(),
) {
	if i > magnitude {
		above()
	} else if i*-1 > magnitude {
		below()
	}
}
