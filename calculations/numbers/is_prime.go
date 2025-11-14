package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	if count == 2 {
		return true
	}
	return false
}
