package pkg

import "crypto/sha256"

const base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

func byteSliceToInt64(b []byte) int64 {
	var n int64
	for i := range b {
		n = (n << 8) | int64(b[i])
	}

	return n
}

func int64ToBase62(n int64) string {
	if n == 0 {
		return string(base62[0])
	}
	s := ""
	for n > 0 {
		s = string(base62[n%62]) + s
		n /= 62
	}

	return s
}

func GenerateURLFingerprint(longUrl string) string {
	hasher := sha256.New()
	hasher.Write([]byte(longUrl))
	hashBytes := hasher.Sum(nil)
	bytesToEncode := hashBytes[:8]

	numericRepresentation := byteSliceToInt64(bytesToEncode)
	shortUrl := int64ToBase62(numericRepresentation)

	return shortUrl
}
