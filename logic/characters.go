package logic

func Characters(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 32 || s[i] > 126) && s[i] != 10 {
			return false
		}
	}
	return true
}
