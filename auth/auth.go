package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/seuc-frp-utn/api/roles"
	"os"
)

var (
	secretKey []byte
	audience string
)

type Profile struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Roles roles.Role
}

type JWT struct {
	UUID string
	Name string
	Email string
	Roles roles.Role
}

type Claims struct {
	Profile  Profile `json:"profile,omitempty"`
	jwt.StandardClaims
}


func init() {
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	audience = os.Getenv("JWT_AUDIENCE")
}

func Encode(body JWT) (*string, error) {
	claims := Claims{
		Profile{
			Email: body.Email,
			Name:  body.Name,
			Roles: body.Roles,
		},
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "seuc.frp.utn.edu.ar",
			Subject:   body.UUID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

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
}

func Sanitize(token string) bool {
	var header string
	var payload string
	var signature string
	fmt.Sscanf(token, "%s.%s.%s", &header, &payload, &signature)
	if len(header) > 0 && len(payload) > 0 && len(signature) > 0 {
		return true
	}
	return false
}