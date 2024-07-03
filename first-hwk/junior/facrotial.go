package junior

func GetFactorial(a uint64) uint64 {
	if a == 0 {
		return 1
	}
	return a * GetFactorial(a-1)
}
