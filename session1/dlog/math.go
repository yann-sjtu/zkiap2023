package dlog

func index2(x int) []int {
	var exponents []int
	for x != 0 {
		exponents = append(exponents, x&1)
		x /= 2
	}
	return exponents
}

func exp(g, x, p int) int {
	if x < 5 {
		result := 1
		for i := 0; i < x; i++ {
			result = result * g % p
		}
		return result
	}

	exponents := index2(x)
	result := 1
	if exponents[0] == 1 {
		result = g
	}
	mul := g * g % p
	for e, o := range exponents[1:] {
		if o == 0 {
			continue
		}
		n := mul
		for i := 0; i < e; i++ {
			n = n * n % p
		}

		result = result * n % p
	}
	return result
}
