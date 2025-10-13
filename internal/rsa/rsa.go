package rsa

import (
	"backend/internal/alphabet"
	"fmt"
)

type RSA struct {
	P, Q   int
	N, Phi int
	D, E   int
}

func NewRSA(maxPrime int) (*RSA, error) {
	p, q, n, phi, d, e, err := RandomKeys(maxPrime)
	if err != nil {
		return nil, err
	}
	return &RSA{P: p, Q: q, N: n, Phi: phi, D: d, E: e}, nil
}

func NewRSAManual(p, q, d int) (*RSA, error) {
	if p == q {
		return nil, fmt.Errorf("p и q должны быть разными")
	}

	n := p * q
	phi := (p - 1) * (q - 1)

	if Euclid(d, phi) != 1 {
		return nil, fmt.Errorf("d и phi не взаимно просты")
	}

	// Вычисляем e по формуле расширенного Евклида
	e, err := OpenKey(d, phi)
	if err != nil {
		return nil, err
	}

	return &RSA{
		P: p, Q: q, N: n, Phi: phi, D: d, E: e, // <- D оставляем таким, каким передали
	}, nil
}

func (r *RSA) Cipher(text string) []int {
	result := []int{}
	for _, ch := range text {
		if val, ok := alphabet.Alphabet[ch]; ok {
			result = append(result, FastPowMod(val, r.E, r.N))
		}
	}
	return result
}

func (r *RSA) Decipher(cipher []int) string {
	text := ""
	for _, val := range cipher {
		num := FastPowMod(val, r.D, r.N)
		ch, ok := alphabet.ReversedAlphabet[num]
		if !ok {
			fmt.Println("Не найдено в алфавите:", num)
		} else {
			text += ch
		}
	}
	fmt.Println("Расшифрованный текст:", text)
	return text
}
