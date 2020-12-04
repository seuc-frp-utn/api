package auth

import "golang.org/x/crypto/bcrypt"

// GeneratePassword generates a bcrypt password from the given plain password.
func GeneratePassword(plain string) (*string, error) {
	bslice := []byte(plain)

	generated, err := bcrypt.GenerateFromPassword(bslice, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashed := string(generated)

	return &hashed, nil
}

// ComparePasswords receives the plain and the hashed password,
// encodes the first one and, tries to compare both of them.
// It returns true if they are equal. Returns false in any other case.
func ComparePasswords(plain, hashed string) bool {
	bPlain := []byte(plain)
	bHashed := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(bHashed, bPlain)
	if err != nil {
		return false
	}
	return true
}
