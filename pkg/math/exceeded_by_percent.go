package math

func ExceededByPercent(
	past float64,
	now float64,
	hundredPercent float64,
	percent float64,
) bool {
	if now <= hundredPercent {
		return false
	}

	percentAmount := hundredPercent * percent / 100

	if now < hundredPercent+percentAmount {
		return false
	}

	pastDivided := past / percentAmount
	nowDivided := now / percentAmount

	return int(nowDivided) > int(pastDivided)
}
