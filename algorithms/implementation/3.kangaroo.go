func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	var meet string
	var kang1 int32 = x1
	var kang2 int32 = x2
	speed1 := v1
	speed2 := v2

	for kang2 = x2; kang2 >= kang1; kang2 += speed2 {
		if kang2 > kang1 {
			meet = "NO"
		} else if kang2 == kang1 {
			meet = "YES"
		}
		kang1 += speed1

	}

	return meet
}