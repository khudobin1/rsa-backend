package rsa

import (
	"fmt"
	"math/rand"
	"time"
)

func ClosedKeys(phi int) []int {
	keys := []int{}
	for i := 2; i < phi; i++ {
		if Euclid(i, phi) == 1 {
			keys = append(keys, i)
		}
	}
	return keys
}

func OpenKey(d, phi int) (int, error) {
	gcd, x, _ := ExtendedEuclid(d, phi)
	if gcd != 1 {
		return 0, fmt.Errorf("НОД(d, phi) != 1")
	}
	if x < 0 {
		x += phi
	}
	return x, nil
}

func RandomKeys(maxPrime int) (p, q, n, phi, d, e int, err error) {
	primes := Sieve(maxPrime)
	if len(primes) < 2 {
		return 0, 0, 0, 0, 0, 0, fmt.Errorf("слишком мало простых чисел")
	}

	rand.Seed(time.Now().UnixNano())
	p = primes[rand.Intn(len(primes))]
	q = primes[rand.Intn(len(primes))]
	for q == p {
		q = primes[rand.Intn(len(primes))]
	}

	n = p * q
	phi = (p - 1) * (q - 1)
	closed := ClosedKeys(phi)
	d = closed[rand.Intn(len(closed))]
	e, err = OpenKey(d, phi)
	return
}
