package student

func CheckNumber(arg string) bool {
	list := "0123456789"
	for _, char := range arg {
		for _, nbr := range list {
			if char != nbr {
				continue
			} else {
				return true
			}
		}
	}
	return false
}
