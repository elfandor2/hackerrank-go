func gradingStudents(grades []int32) []int32 {
	// Write your code here
	roundGrades := []int32{}
	for _, v := range grades {
		lastDigit := v % 10
		if v > 37 && (lastDigit == 3 || lastDigit == 8) {
			roundGrades = append(roundGrades, v+2)
		} else if v > 37 && (lastDigit == 4 || lastDigit == 9) {
			roundGrades = append(roundGrades, v+1)
		} else {
			roundGrades = append(roundGrades, v)
		}
	}
	return roundGrades
}