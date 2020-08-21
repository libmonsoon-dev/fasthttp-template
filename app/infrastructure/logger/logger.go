package logger

import (
	"io"
	"log"
	"os"
)

type Logger = log.Logger

func NewStderrLogger() *Logger {
	return New(os.Stderr)
}

func New(out io.Writer) *Logger {
	const flags = log.Llongfile | log.Ltime | log.Ldate
	return log.New(out, "", flags)
}
