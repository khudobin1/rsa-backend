package rsa

import "math"

func Sieve(n int) []int {
	if n < 2 {
		return []int{}
	}
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func Euclid(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func ExtendedEuclid(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := ExtendedEuclid(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

func FastPowMod(a, n, m int) int {
	p := 1
	ak := a % m
	for n > 0 {
		if n%2 == 1 {
			p = (p * ak) % m
		}
		ak = (ak * ak) % m
		n /= 2
	}
	return p
}
