package auth

import guuid "github.com/google/uuid"

// GenerateUUID generates an UUID string.
func GenerateUUID() string {
	uuid := guuid.New()
	return uuid.String()
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