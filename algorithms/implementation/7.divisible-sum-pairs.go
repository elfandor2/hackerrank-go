func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var result int32

	for i := 0; i < int(n)-1; i++ {
		for j := 1; j <= int(n)-i-1; j++ {
			if (ar[i]+ar[i+j])%k == 0 {
				result++
			}
			fmt.Println(ar[i], "+", ar[j])
		}
	}

	return result
}