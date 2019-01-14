package logger

import (
	"io"
	"log"
)

type logger struct {
	*log.Logger
}

type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

func New(out io.Writer, prefix string, flag int) Logger {
	return &logger{log.New(out, prefix, flag)}
}

type T interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

func NewTadapter(t T) Logger {
	return &tAdapter{t}
}

type tAdapter struct {
	T
}

func (a *tAdapter) Print(args ...interface{}) {
	a.Log(args...)
}

func (a *tAdapter) Printf(format string, args ...interface{}) {
	a.Logf(format, args...)
}
