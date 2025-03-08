package utils

import (
	"crypto/sha256"
	"encoding/binary"
)

// GenerateVector converts text into a hashed float32 vector.
func GenerateVector(data string) []float32 {
	hash := sha256.Sum256([]byte(data))
	vector := make([]float32, 3)

	for i := 0; i < 3; i++ {
		vector[i] = float32(binary.BigEndian.Uint32(hash[i*8:(i+1)*8])) / 1e9
	}
	return vector
}
