package app

type Root struct {
	*Config
	UserService
	AuthService
}
