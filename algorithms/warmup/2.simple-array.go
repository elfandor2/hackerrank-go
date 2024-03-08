func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var result int32
	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}
	return result
}