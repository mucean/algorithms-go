package integer

import "math"

func Divide(divided, divisor int32) int32 {
	if divided == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	negative := 2
	if divided > 0 {
		divided = -divided
		negative--
	}
	if divisor > 0 {
		divisor = -divisor
		negative--
	}
	if negative == 1 {
		return -divideCore(divided, divisor)
	}
	return divideCore(divided, divisor)
}

var halfMin int32 = math.MinInt32 / 2

func divideCore(divided, divisor int32) int32 {
	var result int32 = 0

	for divided <= divisor {
		val := divisor
		var quotient int32 = 1
		for val >= halfMin && divided <= val+val {
			quotient += quotient
			val += val
		}
		divided -= val
		result += quotient
	}

	return result
}
