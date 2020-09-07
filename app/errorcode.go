package app

type ErrorCode int

const (
	CodeInternalError ErrorCode = iota + 1
	CodeItemNotFound
	CodePasswordNotMatch
	CodeInvalidParams
)
