
func compareTriplets(a []int32, b []int32) []int32 {
	result := []int32{0, 0}
	for i, _ := range a {
		if a[i] > b[i] {
			result[0]++
		} else if b[i] > a[i] {
			result[1]++
		}
	}
	return result
}

