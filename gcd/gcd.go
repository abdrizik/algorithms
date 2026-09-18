package gcd

// GCD returns the greatest common divisor of a and b, using Euclid's algorithm
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
