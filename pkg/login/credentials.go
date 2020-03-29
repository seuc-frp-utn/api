package login

type CredentialsInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type CredentialsOutput struct {
	Token string
	ExpiresAt int64
	Subject string
}