package utils

// FixKey ensures the key is of valid length (16, 24, or 32 bytes).
func FixKey(key string, length int) string {
	if len(key) > length {
		return key[:length]
	}
	for len(key) < length {
		key += "0"
	}
	return key
}
