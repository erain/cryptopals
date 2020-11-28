package set1

import (
	"testing"

	"github.com/erain/cryptopals/utils"
)

func TestDecodeSingleCharXORCipher(t *testing.T) {
	encoded := utils.HexDecode("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	key := DecodeSingleCharXORCipher(encoded)
	t.Logf("Key is '%v'", string(key))
}
