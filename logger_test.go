package logger

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/gregoryv/asserter"
)

func Test_constructor(t *testing.T) {
	l := New("prefix")
	assert := asserter.New(t)
	assert(l != nil).Fail()
}

func Test_output_of_logger(t *testing.T) {
	buf := bytes.NewBufferString("")
	assert := asserter.New(t)
	cases := []struct {
		l Logger
	}{
		{Adapt(log.New(buf, "", log.LstdFlags))},
	}
	for i, c := range cases {
		exp := fmt.Sprintf("hello %v", i)
		c.l.Log(exp)
		assert().Contains(buf.Bytes(), exp)
		buf.Reset()
		c.l.Logf("%s", exp)
		assert().Contains(buf.Bytes(), exp)
		buf.Reset()
	}
}

func Test_silent(t *testing.T) {
	l := Silent
	l.Log("nada")
	l.Logf("%s", "nada")
	//t.Fail()
}
