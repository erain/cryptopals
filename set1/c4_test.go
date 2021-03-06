package set1

import (
	"testing"

	"github.com/erain/cryptopals/utils"
)

func TestDetectSingleCharXOR(t *testing.T) {
	text := utils.ReadFixturesFile("4.txt")
	key := DetectSingleCharXOR(text)
	t.Logf("Key is '%s'", key)
}
