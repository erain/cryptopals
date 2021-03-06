package set1

func ReaptingKeyXOR(in, key []byte) []byte {
	res := make([]byte, len(in))
	keylen := len(key)
	for i := range in {
		res[i] = in[i] ^ key[i%keylen]
	}
	return res
}
