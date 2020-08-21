package app

type Logger interface {
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}
