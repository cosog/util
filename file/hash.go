// hash_file
package util_hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func HashSHA256File(filePath string) (string, error) {
	var hashValue string
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return hashValue, err
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return hashValue, err
	}
	hashInBytes := hash.Sum(nil)
	hashValue = hex.EncodeToString(hashInBytes)
	return hashValue, nil
}
