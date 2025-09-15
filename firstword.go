package student

func FirstWord(s string) string {
	res := ""
	for _, letter := range s {
		if letter == ' ' {
			break
		}
		res += string(letter)
	}
	return res + "\n"
}
