package string

import (
	"crypto/rand"
	"fmt"
	"time"
)

const (
	letters        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits  = 6 // 6 bits to represent a letter index
	idLen          = 8
	defaultRandLen = 8
	letterIdxMask  = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax   = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// src is a global, locked source of random numbers.
var src = newLockedSource(time.Now().UnixNano())

// RandN returns a random string of length n.
func RandN(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// Rand returns a random string of default length.
func Rand() string {
	return RandN(defaultRandLen)
}

// Seed sets the seed to seed.
func Seed(seed int64) {
	src.Seed(seed)
}

// RandId returns a random id string.
func RandId() string {
	b := make([]byte, idLen)
	_, err := rand.Read(b)
	if err != nil {
		return RandN(idLen)
	}

	return fmt.Sprintf("%x%x%x%x", b[0:2], b[2:4], b[4:6], b[6:8])
}
