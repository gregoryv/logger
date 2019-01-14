package logger

import (
	"bytes"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_output_of_logger(t *testing.T) {
	buf := bytes.NewBufferString("")
	assert := asserter.New(t)
	verbose := NewVerbose(buf)
	debug := NewDebug(buf)
	cases := []struct {
		l L
	}{
		{verbose},
		{debug},
	}
	for _, c := range cases {
		c.l.Log("hello")
		assert().Contains(buf.Bytes(), "hello")
		c.l.Logf("%s is ok", "world")
		assert().Contains(buf.Bytes(), "world is ok")
		buf.Reset()
	}
}

func Test_silent(t *testing.T) {
	l := New()
	l.Log("nada")
	l.Logf("%s", "nada")
	//t.Fail()
}
