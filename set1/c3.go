package set1

import (
	"github.com/erain/cryptopals/utils"
)

// SingleXOR calculate the XOR between an input and a one byte char
func SingleXOR(in []byte, key byte) []byte {
	res := make([]byte, len(in))
	for i, c := range in {
		res[i] = c ^ key
	}
	return res
}

// DecodeSingleCharXORCipher decodes a XOR cipher with one char key
func DecodeSingleCharXORCipher(in []byte) []byte {
	var res []byte
	var lastScore float64

	c := utils.GetDefaultCorpus()

	for key := 0; key < 256; key++ {
		out := SingleXOR(in, byte(key))
		score := utils.ScoreEnglish(string(out), c)
		if score > lastScore {
			res = out
			lastScore = score
		}
	}

	return res
}
