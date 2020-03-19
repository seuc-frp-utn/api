package auth

import (
	"github.com/btcsuite/btcutil/base58"
	guuid "github.com/google/uuid"
	nanoid "github.com/matoous/go-nanoid"
)

// GenerateUUID generates an UUID string.
func GenerateUUID() string {
	uuid := guuid.New()
	return uuid.String()
}

// GenerateShortUUID generates an UUID and encodes it using base58.
func GenerateBase58UUID() string {
	uuid := guuid.New()
	b := []byte(uuid.String())
	return base58.Encode(b)
}

// EncodeUUID receives an UUID and encodes it using base 58.
func EncodeUUID(uuid string) string {
	b := []byte(uuid)
	return base58.Encode(b)
}

// GenerateNanoUUID generates a user-friendly UUID of length 8.
func GenerateNanoUUID() (string, error) {
	uuid, err := nanoid.Generate("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 8)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

// IsUUID receives an UUID value and checks if it's a valid UUID.
func IsUUID(value string) bool {
	uuid, err := guuid.Parse(value)
	if err != nil {
		return false
	}
	if uuid.Version() != 4 {
		return false
	}
	if uuid.Variant() != guuid.RFC4122 {
		return false
	}
	return true
}