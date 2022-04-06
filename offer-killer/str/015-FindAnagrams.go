package str

var FindAnagrams = findAnagrams

func findAnagrams(s1, s2 string) []int {
	l1, l2 := len(s1), len(s2)
	poses := make([]int, 0)
	if l1 > l2 {
		return poses
	}
	byteMap := make([]int, 26)
	for i := range s1 {
		byteMap[s1[i]-'a']++
		byteMap[s2[i]-'a']--
	}
	if allZero(byteMap) {
		poses = append(poses, 0)
	}
	p1, p2 := 0, l1
	for p2 < l2 {
		byteMap[s2[p1]-'a']++
		byteMap[s2[p2]-'a']--
		p1++
		p2++
		if allZero(byteMap) {
			poses = append(poses, p1)
		}
	}
	return poses
}
