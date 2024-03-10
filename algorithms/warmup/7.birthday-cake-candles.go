func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var totalCandle int32
	sort.Slice(candles, func(i, j int) bool {
		return candles[i] > candles[j]
	})
	for _, v := range candles {
		if v == candles[0] {
			totalCandle++
		}
	}
	fmt.Println(candles)
	return totalCandle

}