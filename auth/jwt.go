package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/seuc-frp-utn/api/roles"
	"os"
	"strings"
	"time"
)

var (
	secretKey []byte
	audience string
)

// Profile represents a set of data that is going to be in the JWT body.
type Profile struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Roles roles.Role
}

// JWT represents the JWT body that the Authentication algorithm is going to encode.
type JWT struct {
	UUID string
	Name string
	Email string
	Roles roles.Role
}

// Claims extends the default claims from the JWT package.
type Claims struct {
	Profile  Profile `json:"profile,omitempty"`
	jwt.StandardClaims
}


func init() {
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}

// Encode receives a JWT body and retrieves the access token and the time for it to expire.
// If there is an error, an error is returned instead.
func Encode(body JWT) (*string, *int64, error) {
	expiresAt := time.Now().Add(time.Hour * 8).Unix()
	claims := Claims{
		Profile{
			Email: body.Email,
			Name:  body.Name,
			Roles: body.Roles,
		},
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "seuc.frp.utn.edu.ar",
			Subject:   body.UUID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(secretKey)
	if err != nil {
		return nil, nil, err
	}
	return &result, &expiresAt, nil
}

// Decode tries to convert a string to a JWT.
// It returns an error if something goes wrong.
func Decode(token string) (*JWT, error) {
	parsed, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsed.Claims.(*Claims); ok && parsed.Valid {
		return &JWT{
			UUID:  claims.Subject,
			Name:  claims.Profile.Name,
			Email: claims.Profile.Email,
			Roles: claims.Profile.Roles,
		}, nil
	}
	return nil, errors.New("unable to decode JWT")
}

// SplitJWT receives an string and returns the header, the payload and the signature of a JWT.
func SplitJWT(token string) (header, payload, signature *string) {
	slice := strings.Split(token, ".")
	if len(slice) != 3 {
		return nil, nil, nil
	}
	return &slice[0], &slice[1], &slice[2]
}

// Sanitize makes sure that the given token is valid.
func Sanitize(token string) bool {
	header, payload, signature := SplitJWT(token)
	if header == nil || payload == nil || signature == nil {
		return false
	}
	if len(*header) > 20 && len(*payload) > 20 && len(*signature) > 20 {
		return true
	}
	return false
}