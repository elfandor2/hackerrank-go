func breakingRecords(scores []int32) []int32 {
	// Write your code here
	result := []int32{0, 0}
	var firstScore int32 = scores[0]
	var scoreMax int32 = firstScore
	var scoreMin int32 = firstScore

	for _, v := range scores {
		if v > scoreMax {
			scoreMax = v
			result[0]++
		} else if v < scoreMin {
			scoreMin = v
			result[1]++
		}
	}
	return result
}