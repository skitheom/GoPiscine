package piscine

func swap(a *int, b *int) {

	tmp := *a
	*a = *b
	*b = tmp
}

func countLen(table []int) int {

	length := 0
	for range table {
		length++
	}
	return length
}

func SortIntegerTable(table []int) {

	isSwapped := false
	n := countLen(table)
	for i := 0; i < n-1; i++ {
		isSwapped = false
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				swap(&table[j], &table[j+1])
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
}
