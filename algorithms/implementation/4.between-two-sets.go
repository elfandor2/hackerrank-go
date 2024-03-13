func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	result := 0
	lastA := a[len(a)-1]
	firstB := b[0]
	var numberBetween []int32
	var numberBetween2 []int32

	for i := lastA; i <= firstB; i++ {
		if firstB%i == 0 {
			numberBetween = append(numberBetween, i)
		}
	}

	for _, v := range numberBetween {
		if v%a[0] == 0 && v%a[1] == 0 {
			numberBetween2 = append(numberBetween2, v)
		}
	}
	result = len(numberBetween2)
	return int32(result)
}