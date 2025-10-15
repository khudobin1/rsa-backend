package signature

import (
	"backend/internal/alphabet"
	"math/big"
)

func Hash(n int, H int, text string) int {
	hash := big.NewInt(int64(H))
	mod := big.NewInt(int64(n))

	for _, char := range text {
		c := big.NewInt(int64(alphabet.Alphabet[char]))
		c.Add(c, hash)
		hash.Exp(c, big.NewInt(2), mod)
	}

	return int(hash.Int64())
}

func Signature(hash, d, e, n int) (signatureValue, verifiedHash *big.Int) {
	hashBig := big.NewInt(int64(hash))
	eBig := big.NewInt(int64(e))
	dBig := big.NewInt(int64(d))
	nBig := big.NewInt(int64(n))

	signatureValue = new(big.Int).Exp(hashBig, eBig, nBig)

	verifiedHash = new(big.Int).Exp(signatureValue, dBig, nBig)
	return
}
