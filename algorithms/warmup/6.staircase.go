func staircase(n int32) {
	// Write your code here
	for i := 0; i < int(n); i++ {
		space := " "
		repeatedSpace := strings.Repeat(space, int(n))
		hashtag := "#"
		repeatedHashtag := strings.Repeat(hashtag, int(i+1))
		fmt.Println(repeatedSpace[:len(repeatedSpace)-1-i] + repeatedHashtag)
	}
}