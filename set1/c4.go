package set1

import (
	"strings"

	"github.com/erain/cryptopals/utils"
)

// DetectSingleCharXOR solve the challenges in https://cryptopals.com/sets/1/challenges/4
func DetectSingleCharXOR(text string) []byte {
	var lastScore float64
	var res []byte
	for _, line := range strings.Split(text, "\n") {
		out, score := DecodeSingleCharXORCipher(utils.HexDecode(line))
		if score > lastScore {
			res = out
			lastScore = score
		}
	}
	return res
}
