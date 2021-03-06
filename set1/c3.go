package set1

import (
	"github.com/erain/cryptopals/utils"
)

// DecodeSingleCharXORCipher decodes a XOR cipher with one char key
func DecodeSingleCharXORCipher(in []byte) ([]byte, float64) {
	var res []byte
	var lastScore float64

	c := utils.GetDefaultCorpus()

	for key := 0; key < 256; key++ {
		out := utils.SingleXOR(in, byte(key))
		score := utils.ScoreEnglish(string(out), c)
		if score > lastScore {
			res = out
			lastScore = score
		}
	}

	return res, lastScore
}
