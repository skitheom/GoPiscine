package piscine

func strLen(s string) int {

	l := 0
	for range s {
		l++
	}
	return l
}

func StrRev(s string) string {

	tmp := []rune(s)
	i := strLen(s) - 1
	for _, r := range s {
		tmp[i] = r
		i--
	}
	return string(tmp)
}
