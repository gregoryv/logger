package logger

import (
	"io"
	"log"
)

// Returns a logger with log.LstdFlags
func New(w io.Writer) L {
	return &lAdapter{log.New(w, "", log.LstdFlags)}
}

type Any interface{}

type L interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

type lAdapter struct{ P }

func (a *lAdapter) Log(args ...interface{})            { a.Print(args...) }
func (a *lAdapter) Logf(f string, args ...interface{}) { a.Printf(f, args...) }

func Adapt(printer P) L { return &lAdapter{printer} }

type P interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

// Returns a logger with log.LstdFlags|log.Lshortfile
func NewDebug(w io.Writer) L {
	return &lAdapter{log.New(w, "", log.LstdFlags|log.Lshortfile)}
}

// Logger that outputs nothing
const Silent silent = iota

type silent int

func (silent) Log(args ...interface{})                 {}
func (silent) Logf(format string, args ...interface{}) {}
