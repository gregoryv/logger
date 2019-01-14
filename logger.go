package logger

import (
	"io"
	"log"
)

// Returns a logger with log.LstdFlags
func NewVerbose(w io.Writer) L {
	return &pAdapter{log.New(w, "", log.LstdFlags)}
}

// Returns a logger with log.LstdFlags|log.Lshortfile
func NewDebug(w io.Writer) L {
	return &pAdapter{log.New(w, "", log.LstdFlags|log.Lshortfile)}
}

type L interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

type pAdapter struct {
	*log.Logger
}

func (p *pAdapter) Log(args ...interface{}) {
	p.Print(args...)
}

func (p *pAdapter) Logf(format string, args ...interface{}) {
	p.Printf(format, args...)
}

type P interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

// Returns a silent logger
func New() L {
	return &silent{}
}

type silent struct{}

func (*silent) Log(args ...interface{})                 {}
func (*silent) Logf(format string, args ...interface{}) {}
