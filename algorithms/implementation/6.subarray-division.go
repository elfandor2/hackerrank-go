func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var result int32

	for i := 0; i < len(s)+1-int(m); i++ {
		var subBirthday int32
		for j := 0; j < int(m); j++ {
			subBirthday += s[i+j]
		}
		if subBirthday == d {
			result++
		}
		subBirthday = 0
	}
	return result
}