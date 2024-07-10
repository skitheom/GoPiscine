package piscine

func Atoi(s string) int {

	isNeg := false
	uint64IntMax := uint64(^uint(0) >> 1)
	uint64IntMin := uint64IntMax + 1
	n := uint64(0)
	for i, r := range s {
		if i == 0 && (r == '+' || r == '-') {
			if r == '-' {
				isNeg = true
			}
			continue
		}
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + uint64(r-'0')
		if (!isNeg && n > uint64IntMax) || (isNeg && n > uint64IntMin) {
			return 0
		}
	}
	if isNeg {
		if n == uint64IntMin {
			return -int(uint64IntMax) - 1
		}
		return -int(n)
	}
	return int(n)
}
