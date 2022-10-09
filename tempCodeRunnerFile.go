func divide(a int, b int) int {
	if a == math.MinInt32 { // 考虑被除数为最小值的情况
		if b == 1 {
			return math.MinInt32
		}
		if b == -1 {
			return math.MaxInt32
		}
	}
	if a == math.MaxInt32 { // 考虑被除数为最小值的情况
		if b == 1 {
			return math.MaxInt32
		}
		if b == -1 {
			return math.MinInt32
		}
	}
	if b == math.MinInt32 { // 考虑除数为最小值的情况
		if a == math.MinInt32 {
			return 1
		}
		return 0
	}

	if a == 0 {
		return 0

	}
	if a < 0 || b < 0 {
		if a < 0 && b < 0 {
			a = -a
			b = -b
			return calc(a, b)
		} else if a < 0 && b > 0 {
			a = -a
			res := -1 * calc(a, b)
			if res < math.MinInt32 {
				return math.MinInt32
			}
			return res
		} else {
			b = -b
			res := -1 * calc(a, b)
			if res < math.MinInt32 {
				return math.MinInt32
			}
			return res
		}
	}
	return calc(a, b)
}

func calc(a, b int) int {
	result := 0
	if a < b {
		return 0
	} else {
		result = int(math.Pow(2.0, (math.Log2(float64(a)) - math.Log2(float64(b)))))
		if result >= math.MaxInt32 {
			return math.MaxInt32
		}
		return result
	}
}
