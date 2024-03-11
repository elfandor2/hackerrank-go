func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var appleLands int32
	var orangeLands int32
	for i := 0; i < len(apples); i++ {
		if a+apples[i] >= s && a+apples[i] <= t {
			appleLands++
		}
	}
	for i := 0; i < len(oranges); i++ {
		if b+oranges[i] <= t && b+oranges[i] >= s {
			orangeLands++
		}
	}
	fmt.Println(appleLands)
	fmt.Println(orangeLands)

}