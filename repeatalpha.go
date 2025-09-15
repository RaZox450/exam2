package student

func RepeatAlpha(s string) string {
	alphabet := "abcdefghijklmnopkrstuvwxyz"
	var res string
	for _, char := range s {
		if !contains(alphabet, char) {
			res += string(char)
		} else {
			for i := range alphabet {
				if string(char) == alphabet[i] {
					for j := 0; j <= i; j++ {
						res = string(char)
					}
				}
			}
		}
	}
	return res
}

func contains(s string, letter rune) bool {
	for _, l := range s {
		if l == letter {
			return true
		}
	}
	return false
}
