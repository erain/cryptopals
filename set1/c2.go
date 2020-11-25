package set1

// FixedXOR xors for two equal length string
func FixedXOR(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor: mismatched lengths")
	}
	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}

	return res
}
