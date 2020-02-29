package auth

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(plain string) (*string, error) {
	bslice := []byte(plain)

	generated, err := bcrypt.GenerateFromPassword(bslice, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashed := string(generated)

	return &hashed, nil
}

func ComparePasswords(plain, hashed string) bool {
	bPlain := []byte(plain)
	bHashed := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(bHashed, bPlain)
	if err != nil {
		return false
	}
	return true
}
