func timeConversion(s string) string {
	// Write your code here
	var militaryFormat string
	var amPM = string(s[(len(s) - 2):])

	hour, _ := strconv.Atoi(string(s[:2]))
	hourPm := hour
	fmt.Println(amPM)

	if amPM == "PM" && hourPm == 12 {
		militaryFormat = string(s[:len(s)-2])
	} else if amPM == "PM" && hourPm < 12 {
		hourPm += 12
		hourString := strconv.Itoa(hourPm)
		militaryFormat = hourString + string(s[2:len(s)-2])
	} else if amPM == "AM" && hourPm == 12 {
		hourPm -= 12
		hourString := strconv.Itoa(hourPm)
		militaryFormat = "0" + hourString + string(s[2:len(s)-2])
	} else if amPM == "AM" && hourPm < 12 {
		militaryFormat = string(s[:len(s)-2])
	}
	return militaryFormat
}