package set1

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	b1 := hexDecode(t, "1c0111001f010100061a024b53535009181c")
	b2 := hexDecode(t, "686974207468652062756c6c277320657965")
	expected := hexDecode(t, "746865206b696420646f6e277420706c6179")
	got := FixedXOR(b1, b2)
	if !bytes.Equal(got, expected) {
		t.Errorf("wrong result: %x", got)
	}
}

func hexDecode(t *testing.T, s string) []byte {
	v, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal("failed to decode hex: ", s)
	}
	return v
}
