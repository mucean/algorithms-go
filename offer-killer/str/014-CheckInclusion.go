package str

var CheckInclusion = checkInclusion

func checkInclusion(s1, s2 string) bool {
	l1 := len(s1)
	if l1 == 0 {
		return true
	}
	l2 := len(s2)
	if l1 > l2 {
		return false
	}
	byteMap := make([]int, 26)

	for i := range s1 {
		byteMap[s1[i]-'a']++
		byteMap[s2[i]-'a']--
	}
	if allZero(byteMap) {
		return true
	}

	p1, p2 := 0, l1

	for p2 < l2 {
		byteMap[s2[p2]-'a']--
		byteMap[s2[p1]-'a']++
		if allZero(byteMap) {
			return true
		}
		p1++
		p2++
	}
	return false
}

func allZero(nums []int) bool {
	for _, v := range nums {
		if v != 0 {
			return false
		}
	}
	return true
}
