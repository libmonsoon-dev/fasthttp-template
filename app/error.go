package app

type Error interface {
	error
	Unwrap() error
	PublicError() (code ErrorCode, publicMessage string)
	HttpStatus() int
}
