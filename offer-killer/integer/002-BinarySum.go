package integer

func BinarySum(bigger, smaller string) string {
	bigger = binaryRemoveLeftZero(bigger)
	bl := len(bigger)
	smaller = binaryRemoveLeftZero(smaller)
	if bl == 0 {
		return smaller
	}
	sl := len(smaller)
	if sl == 0 {
		return bigger
	} else if sl > bl {
		bigger, smaller = smaller, bigger
		bl, sl = sl, bl
	}

	res := make([]byte, bl+1)
	var carry byte = '0'
	var now byte
	for sl = sl - 1; sl >= 0; sl-- {
		now, carry = binaryAdd(carry, smaller[sl])
		if carry == '0' {
			now, carry = binaryAdd(now, bigger[bl-1])
		} else {
			now = bigger[bl-1]
		}

		res[bl] = now
		bl--
	}
	if carry == '0' {
		res = res[1:]
		copy(res[:bl], bigger[:bl])
		return string(res)
	}

	for bl = bl - 1; bl >= 0; bl-- {
		now, carry = binaryAdd(carry, bigger[bl])
		res[bl+1] = now
	}

	if carry == '0' {
		return string(res[1:])
	}
	res[0] = '1'
	return string(res)
}

func binaryAdd(a, b byte) (byte, byte) {
	if a == '0' {
		return b, '0'
	} else if b == '0' {
		return '1', '0'
	} else {
		return '0', '1'
	}
}

func binaryRemoveLeftZero(num string) string {
	l := len(num)
	if l == 0 {
		return ""
	}

	var i int

	for i = 0; i < l; i++ {
		if num[i] == '1' {
			break
		}
	}

	if i == l {
		return ""
	}

	return num[i:]
}
