package set1

import (
	"bytes"
	"testing"

	"github.com/erain/cryptopals/utils"
)

func TestFixedXOR(t *testing.T) {
	b1 := utils.HexDecode("1c0111001f010100061a024b53535009181c")
	b2 := utils.HexDecode("686974207468652062756c6c277320657965")
	expected := utils.HexDecode("746865206b696420646f6e277420706c6179")
	got := FixedXOR(b1, b2)
	if !bytes.Equal(got, expected) {
		t.Errorf("wrong result: %x", got)
	}
}
