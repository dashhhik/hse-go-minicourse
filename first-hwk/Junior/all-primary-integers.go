package Junior

func AllPrimaryIntegers(n int) []int {
	if n < 2 {
		return nil
	}

	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	var result []int
	for i := 2; i <= n; i++ {
		if primes[i] {
			result = append(result, i)
		}
	}

	return result
}
