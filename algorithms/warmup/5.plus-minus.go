func plusMinus(arr []int32) {
	// Write your code here
	var count = []float32{0, 0, 0}
	for _, v := range arr {
		lenArr := float32(len(arr))
		if v > 0 {
			count[0] += 1 / lenArr
		} else if v < 0 {
			count[1] += 1 / lenArr
		} else {
			count[2] += 1 / lenArr
		}
	}
	fmt.Printf("%.6f\n", count[0])
	fmt.Printf("%.6f\n", count[1])
	fmt.Printf("%.6f\n", count[2])

}