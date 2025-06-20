package math

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Calculate x^n in log(n)
func Pow(x, n int) int {
	ret := 1
	for n > 0 {
		if n%2 == 1 {
			ret *= x
		}
		x *= x
		n /= 2
	}
	return ret
}

func GCD(a, b int) int {
	if b == 0 {
		return Abs(a)
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return a * (b / GCD(a, b))
}
