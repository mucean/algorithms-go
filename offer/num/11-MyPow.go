package num

var MyPow = myPow

func myPow(x float64, exp int) float64 {
	if isZero(x) {
		return 0
	}
	// exp == 0
	if exp == 0 {
		return 1
	}
	posExp := exp
	if posExp < 0 {
		posExp = -posExp
	}
	ans := positivePow(x, posExp)
	if exp > 0 {
		return ans
	}
	return 1 / ans
}

func positivePow(x float64, exp int) float64 {
	if exp == 0 {
		return 1
	} else if exp == 1 {
		return x
	}
	if exp&1 == 1 {
		return x * positivePow(x, exp-1)
	}
	ans := positivePow(x, exp/2)
	return ans * ans
}

func FloatEqual(x, y float64) bool {
	res := x - y
	return res > -0.00000001 && res < 0.00000001
}

func isZero(x float64) bool {
	return FloatEqual(x, 0.0)
}
