package auth

import guuid "github.com/google/uuid"

func GenerateUUID() string {
	uuid := guuid.New()
	return uuid.String()
}

func IsUUID(token string) bool {
	uuid, err := guuid.Parse(token)
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