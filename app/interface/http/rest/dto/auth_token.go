package dto

type AuthToken struct {
	Token string `json:"token"`
}

func AuthTokenFrom(token string) AuthToken {
	return AuthToken{token}
}
