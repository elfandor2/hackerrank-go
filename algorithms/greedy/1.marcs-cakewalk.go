func marcsCakewalk(calorie []int32) int64 {
	// Write your code here
	var minimMiles int64
	sort.Slice(calorie, func(i, j int) bool {
		return calorie[i] > calorie[j]
	})

	for i, c := range calorie {
		mileCalc := math.Pow(2, float64(i)) * float64(c)
		minimMiles += int64(mileCalc)
	}

	fmt.Println(minimMiles)
	return minimMiles
}